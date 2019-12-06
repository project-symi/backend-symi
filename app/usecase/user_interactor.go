package usecase

import "project-symi-backend/app/domain"

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Users() (user domain.Users, err error) {
	user, err = interactor.UserRepository.FindAll()
	return
}

func (interactor *UserInteractor) User(id int) (user domain.User, err error) {
	user, err = interactor.UserRepository.FindById(id)
	return
}
