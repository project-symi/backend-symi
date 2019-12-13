package usecase

import (
	"errors"
	"project-symi-backend/app/domain"
	"strconv"
	"time"
)

type FeedbackInteractor struct {
	FeedbackRepository FeedbackRepository
	FeelingRepository  FeelingRepository
	CategoryRepository CategoryRepository
	UserRepository     UserRepository
	PointRepository    PointRepository
}

const (
	Good = iota + 1
	Meh
	Sad
)

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

func (interactor *FeedbackInteractor) StoreFeedback(feedback domain.Feedback) (employeeIdAndPoint domain.UserIdAndPoints, err error) {
	storedfeedback := domain.StoredFeedback{}
	storedfeedback.UserId, err = interactor.UserRepository.FindKeyIdByEmployeeId(feedback.EmployeeId)
	storedfeedback.FeelingId, err = interactor.FeelingRepository.FindIdByName(feedback.Feeling)
	storedfeedback.CategoryId, err = interactor.CategoryRepository.FindIdByName(feedback.Category)
	storedfeedback.RecipientId, err = interactor.UserRepository.FindKeyIdByEmployeeId(feedback.RecipientEmployeeId)
	storedfeedback.NewsId = feedback.NewsId
	storedfeedback.FeedbackNote = feedback.FeedbackNote
	expireDate := calculateExpireDate()
	appliedPoint, err := interactor.applyPointByFeeling(storedfeedback, expireDate)
	employeeIdAndPoint.EmployeeId = feedback.EmployeeId
	employeeIdAndPoint.Points = appliedPoint
	return
}

func (interactor *FeedbackInteractor) applyPointByFeeling(storedfeedback domain.StoredFeedback, expireDate string) (appliedPoint int, err error) {
	switch storedfeedback.FeelingId {
	case Good:
		appliedPoint, err = interactor.FeedbackRepository.StoreFeedbackAndUpdatePoints(storedfeedback, expireDate) //TODO: Validate recipientId without Category-Employee
	case Meh, Sad:
		appliedPoint = 0
		err = interactor.FeedbackRepository.StoreFeedback(storedfeedback)
	default:
		err = errors.New("Invalid feelings") //TODO:think error handling
	}
	return
}

func calculateExpireDate() (endOfQuarter string) {
	quarter := (time.Now().Month()-1)/3 + 1
	switch quarter {
	case 1:
		endOfQuarter = strconv.Itoa(time.Now().Year()) + "-3-31"
	case 2:
		endOfQuarter = strconv.Itoa(time.Now().Year()) + "-6-31"
	case 3:
		endOfQuarter = strconv.Itoa(time.Now().Year()) + "-9-30"
	case 4:
		endOfQuarter = strconv.Itoa(time.Now().Year()) + "-12-31"
	}
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
