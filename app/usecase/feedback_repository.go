package usecase

import "project-symi-backend/app/domain"

type FeedbackRepository interface {
	FindAll() (domain.Feedbacks, error)
	FindByFeeling(string) (domain.Feedbacks, error)
	FindByEmployeeId(int) (domain.Feedbacks, error)
	UpdateSeen(string) (int, error)
	StoreFeedback(domain.StoredFeedback) error
}
