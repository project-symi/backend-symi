package usecase

import "project-symi-backend/app/domain"

type FeedbackInteractor struct {
	FeedbackRepository FeedbackRepository
}

func (interactor *FeedbackInteractor) Feedbacks() (feedback domain.Feedbacks, err error) {
	feedback, err = interactor.FeedbackRepository.FindAll()
	return
}
