package usecase

import "project-symi-backend/app/domain"

type PointsRepository interface {
	FindPointsByUserId(int) (domain.Points, error)
}
