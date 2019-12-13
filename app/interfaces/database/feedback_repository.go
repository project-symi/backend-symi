package database

import (
	"project-symi-backend/app/domain"
	"time"
)

type FeedbackRepository struct {
	SqlHandler
}

const (
	Poll = iota + 1
	SubmittedFeedback
	ReceivedPositiveFeedback
)

func (repo *FeedbackRepository) FindAll() (feedbacks domain.Feedbacks, err error) {
	rows, err := repo.Query(`
		SELECT
			u1.employee_id,
			feel.name,
			feed.seen,
			c.name,
			COALESCE(u2.employee_id, ''),
			COALESCE(feed.news_id, 0),
			feed.feedback_note
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
			employeeId   string
			feeling      string
			seen         bool
			category     string
			recipientId  string
			newsId       int
			feedbackNote string
		)
		if err := rows.Scan(
			&employeeId,
			&feeling,
			&seen,
			&category,
			&recipientId,
			&newsId,
			&feedbackNote); err != nil {
			continue
		}
		feedback := domain.Feedback{
			EmployeeId:          employeeId,
			Feeling:             feeling,
			Seen:                seen,
			Category:            category,
			RecipientEmployeeId: recipientId,
			NewsId:              newsId,
			FeedbackNote:        feedbackNote,
		}
		feedbacks = append(feedbacks, feedback)
	}
	return
}

func (repo *FeedbackRepository) FindByFeeling(feelingQuery string) (feedbacks domain.Feedbacks, err error) {
	rows, err := repo.Query(`
		SELECT
			u1.employee_id,
			feel.name,
			feed.seen,
			c.name,
			COALESCE(u2.employee_id, ''),
			COALESCE(feed.news_id, 0),
			feed.feedback_note
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
			employeeId   string
			feeling      string
			seen         bool
			category     string
			recipientId  string
			newsId       int
			feedbackNote string
		)
		if err := rows.Scan(
			&employeeId,
			&feeling,
			&seen,
			&category,
			&recipientId,
			&newsId,
			&feedbackNote); err != nil {
			continue
		}
		feedback := domain.Feedback{
			EmployeeId:          employeeId,
			Feeling:             feeling,
			Seen:                seen,
			Category:            category,
			RecipientEmployeeId: recipientId,
			NewsId:              newsId,
			FeedbackNote:        feedbackNote,
		}
		feedbacks = append(feedbacks, feedback)
	}
	return
}

func (repo *FeedbackRepository) FindByEmployeeId(userId int) (feedbacks domain.Feedbacks, err error) {
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
		feedback := domain.Feedback{
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

func (repo *FeedbackRepository) StoreFeedbackAndUpdatePoints(storedFeedback domain.StoredFeedback, expireDate string) (points int, err error) {
	pointCategory := SubmittedFeedback
	tx, err := repo.Begin()
	if err != nil {
		return
	}
	feedbackId, err := addTxFeedback(tx, storedFeedback)
	if err != nil {
		tx.Rollback()
		return
	}
	err = addTxPointLog(tx, storedFeedback.UserId, pointCategory, feedbackId, expireDate)
	if err != nil {
		tx.Rollback()
		return
	}
	points, err = findTxPointsById(tx, pointCategory)
	err = updateTxUserTotalPoint(tx, storedFeedback.UserId, points, feedbackId)
	if err != nil {
		tx.Rollback()
		return
	}
	if storedFeedback.RecipientId != 0 {
		pointCategory = ReceivedPositiveFeedback
		err = addTxPointLog(tx, storedFeedback.RecipientId, pointCategory, feedbackId, expireDate)
		if err != nil {
			tx.Rollback()
			return
		}
		recipientPoints := 0
		recipientPoints, err = findTxPointsById(tx, pointCategory)
		if err != nil {
			tx.Rollback()
			return
		}
		err = updateTxUserTotalPoint(tx, storedFeedback.RecipientId, recipientPoints, feedbackId)
		if err != nil {
			tx.Rollback()
			return
		}
	}
	tx.Commit()
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

func addTxPointLog(tx Tx, userId int, pointCategorId int, feedbackId int, expireDate string) (err error) {
	_, err = tx.Execute(`
		INSERT INTO point_logs (user_id, point_category_id, feedback_id, expire_date, created_at, modified_at)
		VALUES (?, ?, ?, ?, ?, ?)
	  `, userId, pointCategorId, feedbackId, expireDate, time.Now(), time.Now())
	if err != nil {
		return
	}
	return
}

func updateTxUserTotalPoint(tx Tx, userId int, point int, feedbackId int) (err error) {
	_, err = tx.Execute(`
	UPDATE users u
	JOIN point_logs p ON p.user_id = u.id
	JOIN point_categories pc ON p.point_category_id = pc.id
	SET
	  u.total_points = u.total_points + pc.point,
	  u.modified_at = ?
	WHERE
	  u.id = ?
	AND
	  p.feedback_id = ?
		`, time.Now(), userId, feedbackId)
	if err != nil {
		return
	}
	return
}

func findTxPointsById(tx Tx, id int) (point int, err error) {
	row, err := tx.Query(`
	SELECT
		point
	FROM point_categories
	WHERE deleted = false
	AND id = ?
	`, id)
	defer row.Close()
	if err != nil {
		return
	}
	row.Next()
	if err = row.Scan(
		&point); err != nil {
		return
	}
	return
}
