package controllers

import (
	"project-symi-backend/app/domain"
	"project-symi-backend/app/interfaces/database"
	"project-symi-backend/app/usecase/interactor"
)

type FeedbackController struct {
	Interactor interactor.FeedbackInteractor
}

func NewFeedbackController(sqlHandler database.SqlHandler) *FeedbackController {
	return &FeedbackController{
		Interactor: interactor.FeedbackInteractor{
			FeedbackRepository: &database.FeedbackRepository{
				SqlHandler: sqlHandler,
			},
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *FeedbackController) AllFeedbacks(c Context) {
	feedbacks, err := controller.Interactor.FindAll()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, feedbacks)
}

func (controller *FeedbackController) FeedbacksByFeeling(c Context) {
	feeling := domain.FeelingQuery{}
	if err := c.ShouldBind(&feeling); err != nil {
		c.JSON(400, ValidationError("FeedbacksByFeeling method's query string is invalid ", err))
		return
	}
	feedbacks, err := controller.Interactor.FindByFeeling(feeling.Feeling)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, feedbacks)
}

func (controller *FeedbackController) FeedbacksByEmployeeId(c Context) {
	employeeId := domain.EmployeeIdParam{}
	if err := c.ShouldBindUri(&employeeId); err != nil {
		c.JSON(400, ValidationError("UsersByEmployeeId method's parameter is invalid ", err))
		return
	}
	feedbacks, err := controller.Interactor.FindByEmployeeId(employeeId.EmployeeId)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, feedbacks)
}

func (controller *FeedbackController) PatchSeen(c Context) {
	var ids []int
	if err := c.BindJSON(&ids); err != nil {
		c.JSON(400, ValidationError("PatchSeen method's json parameter is invalid ", err))
		return
	}
	numberOfChanged, err := controller.Interactor.ChangeToSeen(ids)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	if numberOfChanged == 0 {
		c.Status(400)
		return
	}
	c.Status(200)
}
