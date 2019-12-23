package controllers

import (
	"encoding/json"
	"project-symi-backend/app/domain"
	"project-symi-backend/app/interfaces/database"
	"project-symi-backend/app/interfaces/http_interface"
	"project-symi-backend/app/usecase/interactor"
)

type NewsController struct {
	NewsInteractor  interactor.NewsInteractor
	SlackInteractor interactor.SlackInteractor
	HttpHandler     http_interface.HttpHandler
}

const SlackNameNews = "news updated"

func NewNewsController(sqlHandler database.SqlHandler, httpHandler http_interface.HttpHandler) *NewsController {
	return &NewsController{
		NewsInteractor: interactor.NewsInteractor{
			NewsRepository: &database.NewsRepository{
				SqlHandler: sqlHandler,
			},
		},
		SlackInteractor: interactor.SlackInteractor{
			SlackRepository: &database.SlackRepository{
				SqlHandler: sqlHandler,
			},
		},
		HttpHandler: httpHandler,
	}
}

func (controller *NewsController) AllNews(c Context) {
	news, err := controller.NewsInteractor.News()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, news)
}

func (controller *NewsController) DeleteByNewsId(c Context) {
	amountOfDeleted, err := controller.NewsInteractor.Delete(c.Param("newsId"))
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	if amountOfDeleted == 0 {
		c.Status(400) //TODO: create another error
		return
	}
	c.Status(204)
}

func (controller *NewsController) AddNewsItem(c Context) {
	newsItem := domain.NewsPost{}
	if err := c.ShouldBindJSON(&newsItem); err != nil {
		c.JSON(400, ValidationError("AddNewsItem method", err))
		return
	}
	success, err := controller.NewsInteractor.AddNewNews(newsItem)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	if success == false {
		c.Status(400)
		return
	}
	c.Status(201)

	// Send slack message
	slack, err := controller.SlackInteractor.FindSlackInfo(SlackNameNews)
	if err != nil {
		return
	}
	slackBody := domain.SlackBody{}
	slackBody.Text = slack.Text
	bodyJson, _ := json.Marshal(slackBody)

	controller.HttpHandler.NewHttpRequest(
		"POST",
		slack.Url,
		bodyJson,
	)
	if err != nil {
		return
	}
	controller.HttpHandler.SetHeader("Content-Type", "application/json")
	controller.HttpHandler.SetHeader("Authorization", "Bearer "+slack.Token)
	err = controller.HttpHandler.DoRequest()
	if err != nil {
		return
	}
}
