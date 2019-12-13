package usecase

import "project-symi-backend/app/domain"

type FeedbackPointsRepository interface {
	StoreFeedbackAndUpdatePoints(domain.StoredFeedback, string) (int, error)
}
