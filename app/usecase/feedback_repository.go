package usecase

import "project-symi-backend/app/domain"

type FeedbackRepository interface {
	FindAll() (domain.Feedbacks, error)
	FindByFeeling(feeling string) (domain.Feedbacks, error)
}
