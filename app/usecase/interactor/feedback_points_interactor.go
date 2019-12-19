package interactor

import (
	"errors"
	"project-symi-backend/app/domain"
	"project-symi-backend/app/usecase/repository"
	"strconv"
	"time"
)

type FeedbackPointsInteractor struct {
	TransactionRepository repository.TransactionRepository
	FeedbackRepository    repository.FeedbackRepository
	FeelingRepository     repository.FeelingRepository
	CategoryRepository    repository.CategoryRepository
	UserRepository        repository.UserRepository
}

const (
	Good = iota + 1
	Meh
	Sad
)

var now = time.Now

func (interactor *FeedbackPointsInteractor) StoreFeedback(feedback domain.FeedbackStore) (storedInfo domain.StoredInfo, err error) {
	storedfeedback := domain.StoredFeedback{}
	storedfeedback.UserId, err = interactor.UserRepository.FindKeyIdByEmployeeId(feedback.EmployeeId)
	storedfeedback.FeelingId, err = interactor.FeelingRepository.FindIdByName(feedback.Feeling)
	storedfeedback.CategoryId, err = interactor.CategoryRepository.FindIdByName(feedback.Category)
	storedfeedback.RecipientId, storedfeedback.RecipientSlackId, err = interactor.UserRepository.FindKeyIdAndSlackIdByEmployeeId(feedback.RecipientEmployeeId)
	storedfeedback.NewsId = feedback.NewsId
	storedfeedback.FeedbackNote = feedback.FeedbackNote
	expireDate := calculateExpireDate()
	appliedPoints, recipientPoints, err := interactor.applyPointByFeeling(storedfeedback, expireDate)
	storedInfo.EmployeeId = feedback.EmployeeId
	storedInfo.RecipientSlackId = storedfeedback.RecipientSlackId
	storedInfo.RecipientPoints = recipientPoints
	storedInfo.Points = appliedPoints
	return
}

func (interactor *FeedbackPointsInteractor) applyPointByFeeling(storedfeedback domain.StoredFeedback, expireDate string) (appliedPoints int, recipientPoints int, err error) {
	switch storedfeedback.FeelingId {
	case Good:
		appliedPoints, recipientPoints, err = interactor.TransactionRepository.StoreFeedbackAndUpdatePoints(storedfeedback, expireDate) //TODO: Validate recipientId without Category-Employee
	case Meh, Sad:
		appliedPoints = 0
		err = interactor.FeedbackRepository.StoreFeedback(storedfeedback)
	default:
		err = errors.New("Invalid feelings") //TODO:think error handling
	}
	return
}

func calculateExpireDate() (endOfQuarter string) {
	quarter := (now().Month()-1)/3 + 1
	switch quarter {
	case 1:
		endOfQuarter = strconv.Itoa(now().Year()) + "-03-31"
	case 2:
		endOfQuarter = strconv.Itoa(now().Year()) + "-06-31"
	case 3:
		endOfQuarter = strconv.Itoa(now().Year()) + "-09-30"
	case 4:
		endOfQuarter = strconv.Itoa(now().Year()) + "-12-31"
	}
	return
}
