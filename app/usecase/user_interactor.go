package usecase

import (
	"project-symi-backend/app/domain"
	"strings"
)

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Users() (user domain.Users, err error) {
	user, err = interactor.UserRepository.FindAll()
	return
}

func (interactor *UserInteractor) User(employeeId string) (user domain.User, err error) {
	user, err = interactor.UserRepository.FindByEmployeeId(employeeId)
	return
}

func (interactor *UserInteractor) UsersByName(name string) (users domain.Users, err error) {
	nameArray := strings.Split(name, " ")
	users, err = interactor.UserRepository.FilterByName(nameArray)
	return
}

func (interactor *UserInteractor) Delete(employeeId string) (amountOfDeleted int, err error) {
	amountOfDeleted, err = interactor.UserRepository.DeleteByEmployeeId(employeeId)
	return
}
