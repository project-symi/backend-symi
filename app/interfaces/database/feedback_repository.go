package database

import (
	"project-symi-backend/app/domain"
	"time"
)

type FeedbackRepository struct {
	SqlHandler
}

func (repo *FeedbackRepository) FindAll() (feedbacks domain.Feedbacks, err error) {
	rows, err := repo.Query(`
		SELECT
			feed.id,
			u1.employee_id,
			feel.name,
			feed.seen,
			c.name,
			COALESCE(u2.employee_id, ''),
			COALESCE(feed.news_id, 0),
			feed.feedback_note,
			feed.created_at
  		FROM feedbacks feed
  		JOIN categories c on c.id = feed.category_id
	  	JOIN feelings feel on feel.id = feed.feeling_id
	  	JOIN users u1 on u1.id = feed.user_id
  		LEFT JOIN users u2 on u2.id = feed.recipient_id
	  `)
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var (
			feedbackId   int
			employeeId   string
			feeling      string
			seen         bool
			category     string
			recipientId  string
			newsId       int
			feedbackNote string
			created      string
		)
		if err := rows.Scan(
			&feedbackId,
			&employeeId,
			&feeling,
			&seen,
			&category,
			&recipientId,
			&newsId,
			&feedbackNote,
			&created); err != nil {
			continue
		}
		feedback := domain.Feedback{
			FeedbackId:          feedbackId,
			EmployeeId:          employeeId,
			Feeling:             feeling,
			Seen:                seen,
			Category:            category,
			RecipientEmployeeId: recipientId,
			NewsId:              newsId,
			FeedbackNote:        feedbackNote,
			CreatedAt:           created,
		}
		feedbacks = append(feedbacks, feedback)
	}
	return
}

func (repo *FeedbackRepository) FindByFeeling(feelingQuery string) (feedbacks domain.Feedbacks, err error) {
	rows, err := repo.Query(`
		SELECT
			feed.id,
			u1.employee_id,
			feel.name,
			feed.seen,
			c.name,
			COALESCE(u2.employee_id, ''),
			COALESCE(feed.news_id, 0),
			feed.feedback_note,
			feed.created_at
  		FROM feedbacks feed
  		JOIN categories c on c.id = feed.category_id
  		JOIN feelings feel on feel.id = feed.feeling_id
		JOIN users u1 on u1.id = feed.user_id
  		LEFT JOIN users u2 on u2.id = feed.recipient_id
		WHERE feel.name = ?
	  `, feelingQuery)
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var (
			feedbackId   int
			employeeId   string
			feeling      string
			seen         bool
			category     string
			recipientId  string
			newsId       int
			feedbackNote string
			created      string
		)
		if err := rows.Scan(
			&feedbackId,
			&employeeId,
			&feeling,
			&seen,
			&category,
			&recipientId,
			&newsId,
			&feedbackNote,
			&created); err != nil {
			continue
		}
		feedback := domain.Feedback{
			FeedbackId:          feedbackId,
			EmployeeId:          employeeId,
			Feeling:             feeling,
			Seen:                seen,
			Category:            category,
			RecipientEmployeeId: recipientId,
			NewsId:              newsId,
			FeedbackNote:        feedbackNote,
			CreatedAt:           created,
		}
		feedbacks = append(feedbacks, feedback)
	}
	return
}

func (repo *FeedbackRepository) FindByEmployeeId(userId int) (feedbacks domain.FeedbackEmployees, err error) {
	rows, err := repo.Query(`
		SELECT
			feed.id,
			feel.name,
			feed.seen,
			c.name,
			COALESCE(u.name, ''),
			feed.feedback_note,
			feed.created_at
  		FROM feedbacks feed
  		JOIN feelings feel on feel.id = feed.feeling_id
  		JOIN categories c on c.id = feed.category_id
  		LEFT JOIN users u on u.id = feed.recipient_id
		WHERE feed.user_id = ?
	  `, userId)
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var (
			id            int
			feeling       string
			seen          bool
			category      string
			recipientName string
			feedbackNote  string
			createdAt     string
		)
		if err := rows.Scan(
			&id,
			&feeling,
			&seen,
			&category,
			&recipientName,
			&feedbackNote,
			&createdAt); err != nil {
			continue
		}
		feedback := domain.FeedbackEmployee{
			Id:            id,
			Feeling:       feeling,
			Seen:          seen,
			Category:      category,
			RecipientName: recipientName,
			FeedbackNote:  feedbackNote,
			CreatedAT:     createdAt,
		}
		feedbacks = append(feedbacks, feedback)
	}
	return
}

func (repo *FeedbackRepository) UpdateSeen(ids string) (amountOfAffected int, err error) {
	result, err := repo.Execute(`UPDATE feedbacks SET seen = true, modified_at = ? WHERE id IN `+ids, time.Now())
	if err != nil {
		return
	}
	amountOfAffected64, err := result.RowsAffected()
	if err != nil {
		return
	}
	amountOfAffected = int(amountOfAffected64)
	return
}

func (repo *FeedbackRepository) StoreFeedback(feedback domain.StoredFeedback) (err error) {
	_, err = repo.Execute(`
	INSERT INTO feedbacks (user_id, feeling_id, category_id, recipient_id, news_id, feedback_note, created_at, modified_at)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?)
  `, feedback.UserId, feedback.FeelingId, feedback.CategoryId, feedback.RecipientId, feedback.NewsId, feedback.FeedbackNote, time.Now(), time.Now())
	return
}

func addTxFeedback(tx Tx, feedback domain.StoredFeedback) (insertedId int, err error) {
	result, err := tx.Execute(`
		INSERT INTO feedbacks (user_id, feeling_id, category_id, recipient_id, news_id, feedback_note, created_at, modified_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	  `, feedback.UserId, feedback.FeelingId, feedback.CategoryId, feedback.RecipientId, feedback.NewsId, feedback.FeedbackNote, time.Now(), time.Now())
	insertedId64, err := result.LastInsertId()
	if err != nil {
		return
	}
	insertedId = int(insertedId64)
	return
}
