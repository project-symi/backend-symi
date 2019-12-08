package usecase

import "project-symi-backend/app/domain"

type DepartmentRepository interface {
	FindAll() (departments domain.Departments, err error)
}
