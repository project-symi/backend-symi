package database

import "project-symi-backend/app/domain"

type FeedbackRepository struct {
	SqlHandler
}

func (repo *FeedbackRepository) FindAll() (feedbacks domain.Feedbacks, err error) {
	rows, err := repo.Query("SELECT id, feedback FROM feedbacks")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id int
		var feedback string
		if err := rows.Scan(&id, &feedback); err != nil {
			continue
		}
		feedbackQuery := domain.Feedback{
			ID:       id,
			Feedback: feedback,
		}
		feedbacks = append(feedbacks, feedbackQuery)
	}
	return
}
