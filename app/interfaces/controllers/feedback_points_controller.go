package controllers

import (
	"encoding/json"
	"project-symi-backend/app/domain"
	"project-symi-backend/app/interfaces/database"
	"project-symi-backend/app/interfaces/http_interface"
	"project-symi-backend/app/usecase/interactor"
	"strconv"
	"strings"
)

type FeedbackPointsController struct {
	FeedbackPointsInteractor interactor.FeedbackPointsInteractor
	SlackInteractor          interactor.SlackInteractor
	Httphandler              http_interface.HttpHandler
}

const SlackName = "good feedback"

func NewFeedbackPointsController(sqlHandler database.SqlHandler, httpHandler http_interface.HttpHandler) *FeedbackPointsController {
	return &FeedbackPointsController{
		FeedbackPointsInteractor: interactor.FeedbackPointsInteractor{
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
		SlackInteractor: interactor.SlackInteractor{
			SlackRepository: &database.SlackRepository{
				SqlHandler: sqlHandler,
			},
		},
		Httphandler: httpHandler,
	}
}

func (controller *FeedbackPointsController) PostFeedback(c Context) {
	feedback := domain.FeedbackStore{}
	c.BindJSON(&feedback)
	storedInfo, err := controller.FeedbackPointsInteractor.StoreFeedback(feedback)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, storedInfo)

	// Send slack message
	slack, err := controller.SlackInteractor.FindSlackInfo(SlackName)
	if err != nil {
		return
	}
	slackBody := domain.SlackBody{}
	slackBody.Channel = storedInfo.RecipientSlackId
	slackBody.Text = strings.Replace(slack.Text, "<points>", strconv.Itoa(storedInfo.RecipientPoints), 1)
	bodyJson, _ := json.Marshal(slackBody)
	if storedInfo.RecipientPoints > 0 && storedInfo.RecipientSlackId != "" {
		err = controller.Httphandler.NewHttpRequest(
			"POST",
			slack.Url,
			bodyJson,
		)
		if err != nil {
			return
		}
		controller.Httphandler.SetHeader("Content-Type", "application/json")
		controller.Httphandler.SetHeader("Authorization", "Bearer "+slack.Token)
		err = controller.Httphandler.DoRequest()
		if err != nil {
			return
		}
	}
}
