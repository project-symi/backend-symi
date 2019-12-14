package controllers

import (
	"project-symi-backend/app/interfaces/database"
	"project-symi-backend/app/usecase/interactor"
)

type PointsController struct {
	Interactor interactor.PointsInteractor
}

func NewPointsController(sqlHandler database.SqlHandler) *PointsController {
	return &PointsController{
		Interactor: interactor.PointsInteractor{
			PointsRepository: &database.PointsRepository{
				SqlHandler: sqlHandler,
			},
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *PointsController) PointsByEmployeeId(c Context) {
	points, err := controller.Interactor.FindPointsByEmployeeId(c.Param("employeeId"))
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, points)
}
