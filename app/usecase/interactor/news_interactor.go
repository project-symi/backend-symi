package interactor

import (
	"project-symi-backend/app/domain"
	"project-symi-backend/app/usecase/repository"
)

type NewsInteractor struct {
	NewsRepository repository.NewsRepository
}

func (interactor *NewsInteractor) News() (news domain.News, err error) {
	news, err = interactor.NewsRepository.GetAll()
	return
}

func (interactor *NewsInteractor) Delete(newsId string) (amountOfDeleted int, err error) {
	amountOfDeleted, err = interactor.NewsRepository.DeleteByNewsId(newsId)
	return
}

func (interactor *NewsInteractor) AddNewNews(newsItem domain.NewsPost) (success bool, err error) {
	success, err = interactor.NewsRepository.AddNewsItem(newsItem)
	return
}
