package controllers

import (
	"project-symi-backend/app/interfaces/database"
	"project-symi-backend/app/usecase"
)

type PointsController struct {
	Interactor usecase.PointsInteractor
}

func NewPointsController(sqlHandler database.SqlHandler) *PointsController {
	return &PointsController{
		Interactor: usecase.PointsInteractor{
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
