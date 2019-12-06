package usecase

import "project-symi-backend/app/domain"

type UserRepository interface {
	FindAll() (domain.Users, error)
	FindById(id int) (domain.User, error)
}
