package controllers

import (
	"project-symi-backend/app/domain"
	"project-symi-backend/app/interfaces/database"
	"project-symi-backend/app/usecase/interactor"
)

type FeedbackPointsController struct {
	Interactor interactor.FeedbackPointsInteractor
}

func NewFeedbackPointsController(sqlHandler database.SqlHandler) *FeedbackPointsController {
	return &FeedbackPointsController{
		Interactor: interactor.FeedbackPointsInteractor{
			TransactionRepository: &database.TransactionRepository{
				SqlHandler: sqlHandler,
			},
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

func (controller *FeedbackPointsController) PostFeedback(c Context) {
	feedback := domain.Feedback{}
	c.BindJSON(&feedback)
	pointAndEmployeeId, err := controller.Interactor.StoreFeedback(feedback)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, pointAndEmployeeId)
}
