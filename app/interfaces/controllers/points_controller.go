package controllers

import (
	"project-symi-backend/app/domain"
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
	employeeId := domain.EmployeeIdParam{}
	if err := c.ShouldBindUri(&employeeId); err != nil {
		c.JSON(400, ValidationError("PointsByEmployeeId method's parameter is invalid ", err))
		return
	}
	points, err := controller.Interactor.FindPointsByEmployeeId(employeeId.EmployeeId)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, points)
}
