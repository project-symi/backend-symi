package controllers

import (
	"project-symi-backend/app/interfaces/database"
	"project-symi-backend/app/usecase"
)

type PointController struct {
	Interactor usecase.PointInteractor
}

func NewPointController(sqlHandler database.SqlHandler) *PointController {
	return &PointController{
		Interactor: usecase.PointInteractor{
			PointRepository: &database.PointRepository{
				SqlHandler: sqlHandler,
			},
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *PointController) PointsByEmployeeId(c Context) {
	points, err := controller.Interactor.FindPointsByEmployeeId(c.Param("employeeId"))
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, points)
}
