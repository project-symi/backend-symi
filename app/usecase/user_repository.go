package usecase

import "project-symi-backend/app/domain"

type UserRepository interface {
	FindAll() (domain.Users, error)
	FindByEmployeeId(employeeId string) (domain.User, error)
}
