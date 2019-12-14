package interactor

import (
	"project-symi-backend/app/domain"
	"project-symi-backend/app/usecase/repository"
	"strings"
)

type FeedbackInteractor struct {
	FeedbackRepository repository.FeedbackRepository
	UserRepository     repository.UserRepository
}

func (interactor *FeedbackInteractor) FindAll() (feedback domain.Feedbacks, err error) {
	feedback, err = interactor.FeedbackRepository.FindAll()
	return
}

func (interactor *FeedbackInteractor) FindByFeeling(feeling string) (feedback domain.Feedbacks, err error) {
	feedback, err = interactor.FeedbackRepository.FindByFeeling(feeling)
	return
}

func (interactor *FeedbackInteractor) FindByEmployeeId(employeeId string) (feedback domain.FeedbackEmployees, err error) {
	userId, err := interactor.UserRepository.FindKeyIdByEmployeeId(employeeId)
	feedback, err = interactor.FeedbackRepository.FindByEmployeeId(userId)
	return
}

func (interactor *FeedbackInteractor) ChangeToSeen(ids []string) (numberOfChanged int, err error) {
	query := "(" + strings.Join(ids, ", ") + ")"
	numberOfChanged, err = interactor.FeedbackRepository.UpdateSeen(query)
	return
}
