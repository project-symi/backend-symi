package usecase

import "project-symi-backend/app/domain"

type FeedbackRepository interface {
	FindAll() (domain.Feedbacks, error)
	FindByFeeling(feeling string) (domain.Feedbacks, error)
	InsertFeedback(userId int, feelingId int, categoryId int, recipientId int, newsId int, feedbackNote string) (success bool, err error)
}
