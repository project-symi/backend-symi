package repository

import "project-symi-backend/app/domain"

type FeedbackRepository interface {
	FindAll() (domain.FeedbacksForCEO, error)
	FindByFeeling(string) (domain.Feedbacks, error)
	FindByEmployeeId(int) (domain.FeedbackEmployees, error)
	UpdateSeen(string) (int, error)
	StoreFeedback(domain.StoredFeedback) error
}
