package database

import "project-symi-backend/app/domain"

type FeedbackRepository struct {
	SqlHandler
}

func (repo *FeedbackRepository) FindAll() (feedbacks domain.Feedbacks, err error) {
	rows, err := repo.Query(`
		SELECT
			feel.name,
			feed.seen,
			c.name,
			COALESCE(u.employee_id, ''),
			COALESCE(feed.news_id, 0),
			feed.feedback_note
  		FROM feedbacks feed
  		JOIN categories c on c.id = feed.category_id
	  	JOIN feelings feel on feel.id = feed.feeling_id
  		LEFT JOIN users u on u.id = feed.recipient_id
	  `)
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var (
			feeling      string
			seen         bool
			category     string
			recipientId  string
			newsId       int
			feedbackNote string
		)
		if err := rows.Scan(
			&feeling,
			&seen,
			&category,
			&recipientId,
			&newsId,
			&feedbackNote); err != nil {
			continue
		}
		feedback := domain.Feedback{
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
