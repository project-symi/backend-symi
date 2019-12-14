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
		},
	}
}

func (controller *InvitationController) AllInvitations(c Context) {
	Invitations, err := controller.Interactor.FindBySenderId(c.Query("senderId"))
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, Invitations)
}
