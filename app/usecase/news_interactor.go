package usecase

import (
	"project-symi-backend/app/domain"
)

type NewsInteractor struct {
	NewsRepository NewsRepository
}

func (interactor *NewsInteractor) News() (news domain.News, err error) {
	news, err = interactor.NewsRepository.GetAll()
	return
}

func (interactor *NewsInteractor) Delete(newsId string) (amountOfDeleted int, err error) {
	amountOfDeleted, err = interactor.NewsRepository.DeleteByNewsId(newsId)
	return
}

func (interactor *NewsInteractor) AddNewNews(newsItem domain.NewsItem) (success bool, err error) {
	success, err = interactor.NewsRepository.AddNewsItem(newsItem.Title, newsItem.Description, newsItem.PhotoLink)
	return
}
