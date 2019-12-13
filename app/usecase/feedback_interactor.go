package usecase

import (
	"project-symi-backend/app/domain"
)

type FeedbackInteractor struct {
	FeedbackRepository FeedbackRepository
	UserRepository     UserRepository
}

func (interactor *FeedbackInteractor) FindAll() (feedback domain.Feedbacks, err error) {
	feedback, err = interactor.FeedbackRepository.FindAll()
	return
}

func (interactor *FeedbackInteractor) FindByFeeling(feeling string) (feedback domain.Feedbacks, err error) {
	feedback, err = interactor.FeedbackRepository.FindByFeeling(feeling)
	return
}

func (interactor *FeedbackInteractor) FindByEmployeeId(employeeId string) (feedback domain.Feedbacks, err error) {
	userId, err := interactor.UserRepository.FindKeyIdByEmployeeId(employeeId)
	feedback, err = interactor.FeedbackRepository.FindByEmployeeId(userId)
	return
}

func (interactor *FeedbackInteractor) ChangeToSeen(ids []string) (numberOfChanged int, err error) {
	query := createIdsQuery(ids)
	numberOfChanged, err = interactor.FeedbackRepository.UpdateSeen(query)
	return
}

func createIdsQuery(ids []string) (query string) {
	query = "("
	for i, id := range ids {
		if i != len(ids)-1 {
			query += id + ", "
		} else {
			query += id + ")"
		}
	}
	return
}
