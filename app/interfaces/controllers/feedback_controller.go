package controllers

import (
	"project-symi-backend/app/domain"
	"project-symi-backend/app/interfaces/database"
	"project-symi-backend/app/usecase"
)

type FeedbackController struct {
	Interactor usecase.FeedbackInteractor
}

func NewFeedbackController(sqlHandler database.SqlHandler) *FeedbackController {
	return &FeedbackController{
		Interactor: usecase.FeedbackInteractor{
			FeedbackRepository: &database.FeedbackRepository{
				SqlHandler: sqlHandler,
			},
			FeelingRepository: &database.FeelingRepository{
				SqlHandler: sqlHandler,
			},
			CategoryRepository: &database.CategoryRepository{
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
	feedbacks, err := controller.Interactor.FindByFeeling(c.Query("feeling"))
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, feedbacks)
}

func (controller *FeedbackController) FeedbacksByEmployeeId(c Context) {
	feedbacks, err := controller.Interactor.FindByEmployeeId(c.Param("employeeId"))
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, feedbacks)
}

func (controller *FeedbackController) PostFeedback(c Context) {
	feedback := domain.Feedback{}
	c.BindJSON(&feedback)
	success, err := controller.Interactor.StoreFeedback(feedback)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	if success == false {
		c.Status(200)
		return
	}
	c.Status(201)
}
