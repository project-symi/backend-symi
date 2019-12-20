package interactor

import (
	"project-symi-backend/app/domain"
	"project-symi-backend/app/usecase/repository"
	"strconv"
	"strings"
)

type FeedbackInteractor struct {
	FeedbackRepository repository.FeedbackRepository
	UserRepository     repository.UserRepository
}

func (interactor *FeedbackInteractor) FindAll() (feedback domain.FeedbacksForCEO, err error) {
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

func (interactor *FeedbackInteractor) ChangeToSeen(ids []int) (numberOfChanged int, err error) {
	query := "(" + splitToString(ids, ", ") + ")"
	numberOfChanged, err = interactor.FeedbackRepository.UpdateSeen(query)
	return
}

func splitToString(ints []int, separator string) string {
	if len(ints) == 0 {
		return ""
	}

	stringArr := make([]string, len(ints))
	for i, v := range ints {
		stringArr[i] = strconv.Itoa(v)
	}
	return strings.Join(stringArr, separator)
}
