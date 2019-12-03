package controllers

import (
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
		},
	}
}

func (controller *FeedbackController) Index(c Context) {
	feedbacks, err := controller.Interactor.Feedbacks()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, feedbacks)
}
