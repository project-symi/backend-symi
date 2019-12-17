package controllers

import (
	"project-symi-backend/app/domain"
	"project-symi-backend/app/interfaces/database"
	"project-symi-backend/app/usecase/interactor"
	"strconv"
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
			InvitationStatusCategoryRepository: &database.InvitationStatusCategoryRepository{
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

func (controller *InvitationController) PostInvitation(c Context) {
	invitation := domain.PostInvitation{}
	c.BindJSON(&invitation)
	Invitations, err := controller.Interactor.PostInvitation(invitation)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, Invitations)
}

func (controller *InvitationController) PatchInvitationById(c Context) {
	invitationIdString := c.Param("invitationId")
	invitationId, _ := strconv.Atoi(invitationIdString)
	patchInvitation := domain.PatchInvitation{}
	c.BindJSON(&patchInvitation)
	success, err := controller.Interactor.PatchInvitationById(invitationId, patchInvitation)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	if success == false {
		c.Status(400)
		return
	}
	c.Status(200)
}
