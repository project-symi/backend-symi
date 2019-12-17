package controllers

import (
	"project-symi-backend/app/interfaces/database"
	"project-symi-backend/app/usecase/interactor"
)

type InvitationController struct {
	Interactor interactor.InvitationInteractor
}

func NewInvitationController(sqlHandler database.SqlHandler) *InvitationController {
	return &InvitationController{
		Interactor: interactor.InvitationInteractor{
			InvitationRepository: &database.InvitationRepository{
				SqlHandler: sqlHandler,
			},
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *InvitationController) MadeSeenAllInvitations(c Context) {
	Invitations, err := controller.Interactor.ChangeSeenAndFindAll()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, Invitations)
}

func (controller *InvitationController) InvitationsByEmployeeId(c Context) {
	Invitations, err := controller.Interactor.FindByEmployeeId(c.Param("employeeId"))
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, Invitations)
}
