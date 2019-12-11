package usecase

import "project-symi-backend/app/domain"

type PointInteractor struct {
	PointRepository PointRepository
	UserRepository  UserRepository
}

func (interactor *PointInteractor) FindPointsByEmployeeId(employeeId string) (points domain.Points, err error) {
	userId, err := interactor.UserRepository.FindKeyIdByEmployeeId(employeeId)
	points, err = interactor.PointRepository.FindPointsByUserId(userId)
	if err != nil {
		return
	}
	return
}
