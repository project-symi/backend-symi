package usecase

import "project-symi-backend/app/domain"

type FeedbackRepository interface {
	FindAll() (domain.Feedbacks, error)
	FindByFeeling(string) (domain.Feedbacks, error)
	FindByEmployeeId(int) (domain.Feedbacks, error)
	InsertFeedback(int, int, int, int, int, string) (bool, error)
	UpdateSeen(string) (int, error)
}
