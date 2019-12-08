package usecase

import "project-symi-backend/app/domain"

type UserRepository interface {
	FindAll() (domain.Users, error)
	FindByEmployeeId(employeeId string) (domain.User, error)
	FilterByName(nameArray []string) (domain.Users, error)
	DeleteByEmployeeId(employeeId string) (amountOfDeleted int, err error)
	StoreUsers(query string) (amountOfStored int, err error)
	IsUser(employeeId string) (isUser bool, err error)
}
