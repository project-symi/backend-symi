package usecase

import "project-symi-backend/app/domain"

type TransactionRepository interface {
	StoreFeedbackAndUpdatePoints(domain.StoredFeedback, string) (int, error)
}
