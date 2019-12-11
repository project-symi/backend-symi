package usecase

import "project-symi-backend/app/domain"

type PointRepository interface {
	FindPointsByUserId(int) (domain.Points, error)
}
