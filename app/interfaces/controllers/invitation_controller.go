package controllers

import (
	"project-symi-backend/app/domain"
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
	employeeId := domain.EmployeeIdParam{}
	if err := c.ShouldBindUri(&employeeId); err != nil {
		c.JSON(400, ValidationError("InvitationsByEmployeeId method's parameter is invalid ", err))
		return
	}
	Invitations, err := controller.Interactor.FindByEmployeeId(employeeId.EmployeeId)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, Invitations)
}

func (controller *InvitationController) PostInvitation(c Context) {
	invitation := domain.PostInvitation{}
	if err := c.BindJSON(&invitation); err != nil {
		c.JSON(400, ValidationError("PostInvitation method's json parameter is invalid ", err))
		return
	}
	Invitations, err := controller.Interactor.PostInvitation(invitation)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, Invitations)
}

func (controller *InvitationController) PatchInvitationById(c Context) {
	invitationId := domain.InvitationIdParam{}
	if err := c.ShouldBindUri(&invitationId); err != nil {
		c.JSON(400, ValidationError("InvitationsByEmployeeId method's parameter is invalid ", err))
		return
	}
	patchInvitation := domain.PatchInvitation{}
	if err := c.BindJSON(&patchInvitation); err != nil {
		c.JSON(400, ValidationError("InvitationsByEmployeeId method's json parameter is invalid ", err))
		return
	}
	success, err := controller.Interactor.PatchInvitationById(invitationId.InvitationId, patchInvitation)
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

func (controller *InvitationController) DeleteById(c Context) {
	invitationId := domain.InvitationIdParam{}
	if err := c.ShouldBindUri(&invitationId); err != nil {
		c.JSON(400, ValidationError("DeleteById method's parameter is invalid ", err))
		return
	}
	success, err := controller.Interactor.DeleteById(invitationId.InvitationId)
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
