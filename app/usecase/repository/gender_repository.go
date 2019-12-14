package repository

import "project-symi-backend/app/domain"

type GenderRepository interface {
	FindAll() (genders domain.Genders, err error)
}
