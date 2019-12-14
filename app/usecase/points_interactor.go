package usecase

import "project-symi-backend/app/domain"

type PointsInteractor struct {
	PointsRepository PointsRepository
	UserRepository   UserRepository
}

func (interactor *PointsInteractor) FindPointsByEmployeeId(employeeId string) (points domain.Points, err error) {
	userId, err := interactor.UserRepository.FindKeyIdByEmployeeId(employeeId)
	points, err = interactor.PointsRepository.FindPointsByUserId(userId)
	if err != nil {
		return
	}
	return
}
