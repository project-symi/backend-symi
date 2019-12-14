package interactor

import (
	"project-symi-backend/app/domain"
	"project-symi-backend/app/usecase/repository"
)

type PointsInteractor struct {
	PointsRepository repository.PointsRepository
	UserRepository   repository.UserRepository
}

func (interactor *PointsInteractor) FindPointsByEmployeeId(employeeId string) (points domain.Points, err error) {
	userId, err := interactor.UserRepository.FindKeyIdByEmployeeId(employeeId)
	points, err = interactor.PointsRepository.FindPointsByUserId(userId)
	if err != nil {
		return
	}
	return
}
