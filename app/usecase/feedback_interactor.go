package usecase

import "project-symi-backend/app/domain"

type FeedbackInteractor struct {
	FeedbackRepository FeedbackRepository
	FeelingRepository  FeelingRepository
	CategoryRepository CategoryRepository
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
func (interactor *FeedbackInteractor) StoreFeedback(feedback domain.Feedback) (success bool, err error) {
	userId, err := interactor.UserRepository.FindKeyIdByEmployeeId(feedback.EmployeeId)
	feelingId, err := interactor.FeelingRepository.FindIdByName(feedback.Feeling)
	categoryId, err := interactor.CategoryRepository.FindIdByName(feedback.Category)
	recipientId, err := interactor.UserRepository.FindKeyIdByEmployeeId(feedback.RecipientEmployeeId)
	success, err = interactor.FeedbackRepository.InsertFeedback(userId, feelingId, categoryId, recipientId, feedback.NewsId, feedback.FeedbackNote)
	return
}
