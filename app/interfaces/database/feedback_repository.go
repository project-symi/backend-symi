package database

import "project-symi-backend/app/domain"

type FeedbackRepository struct {
	SqlHandler
}

func (repo *FeedbackRepository) FindAll() (feedbacks domain.Feedbacks, err error) {
	rows, err := repo.Query("SELECT user_id, feedback_note FROM feedbacks")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var userId int
		var feedback string
		if err := rows.Scan(&userId, &feedback); err != nil {
			continue
		}
		feedbackData := domain.Feedback{
			UserId:       userId,
			FeedbackNote: feedback,
		}
		feedbacks = append(feedbacks, feedbackData)
	}
	return
}
