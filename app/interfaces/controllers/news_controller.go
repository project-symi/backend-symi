package controllers

import (
	"project-symi-backend/app/domain"
	"project-symi-backend/app/interfaces/database"
	"project-symi-backend/app/usecase"
)

type NewsController struct {
	Interactor usecase.NewsInteractor
}

func NewNewsController(sqlHandler database.SqlHandler) *NewsController {
	return &NewsController{
		Interactor: usecase.NewsInteractor{
			NewsRepository: &database.NewsRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *NewsController) AllNews(c Context) {
	news, err := controller.Interactor.News()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, news)
}

func (controller *NewsController) DeleteByNewsId(c Context) {
	amountOfDeleted, err := controller.Interactor.Delete(c.Param("newsId"))
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	if amountOfDeleted == 0 {
		c.JSON(400, NewError(err)) //TODO: create another error
		return
	}
	c.Status(204)
}

func (controller *NewsController) AddNewsItem(c Context) {
	newsItem := domain.NewsItem{}
	c.BindJSON(&newsItem)
	success, err := controller.Interactor.AddNewNews(newsItem)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	if success == false {
		c.Status(400)
		return
	}
	c.Status(201)
}
