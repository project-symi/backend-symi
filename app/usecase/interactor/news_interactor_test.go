package interactor_test

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"project-symi-backend/app/domain"
	"project-symi-backend/app/usecase/interactor"
	"project-symi-backend/app/usecase/mock_repository"
	"testing"
)

func SetupNewsTest(t *testing.T) (mockObj *mock_repository.MockNewsRepository) {
	mockCtrl := gomock.NewController(t)
	mockObj = mock_repository.NewMockNewsRepository(mockCtrl)
	defer mockCtrl.Finish()
	return
}

func TestNews(t *testing.T) {
	mockObj := SetupNewsTest(t)

	func() {
		newsItem := domain.NewsItem{
			NewsItemId:  1,
			Title:       "test",
			Description: "This is test news",
			PhotoLink:   "test@test.com",
			Hidden:      false,
			CreatedAt:   "1989-03-15",
			ModifiedAt:  "1989-03-15",
		}
		news := make(domain.News, 1)
		news = append(news, newsItem)
		mockObj.EXPECT().GetAll().Return(news, nil)
		newsInteractorMock := interactor.NewsInteractor{mockObj}
		result, errResult := newsInteractorMock.News()
		if &news == &result || errResult != nil {
			t.Errorf("Cannot get all the news")
		}
	}()

	func() {
		mockObj.EXPECT().GetAll().Return(nil, fmt.Errorf("%s", "getAllQueryError"))
		newsInteractorMock := interactor.NewsInteractor{mockObj}
		result, errResult := newsInteractorMock.News()
		if result != nil || errResult == nil {
			t.Errorf("Cannot get error from GetAll news")
		}

	}()
}

func TestDeleteNewsId(t *testing.T) {
	mockObj := SetupNewsTest(t)

	func() {
		newsId := "1"
		amountOfDeleted := 1
		mockObj.EXPECT().DeleteByNewsId(newsId).Return(amountOfDeleted, nil)
		newsInteractorMock := interactor.NewsInteractor{mockObj}
		result, errResult := newsInteractorMock.Delete(newsId)
		if result != amountOfDeleted || errResult != nil {
			t.Errorf("Cannot delete the news id=%q", newsId)
		}
	}()

	func() {
		newsId := "2"
		amountOfDeleted := 0
		mockObj.EXPECT().DeleteByNewsId(newsId).Return(amountOfDeleted, nil)
		newsInteractorMock := interactor.NewsInteractor{mockObj}
		result, errResult := newsInteractorMock.Delete(newsId)
		if result != amountOfDeleted || errResult != nil {
			t.Errorf("Cannot delete the news id=%q", newsId)
		}
	}()

	func() {
		newsId := "3"
		mockObj.EXPECT().DeleteByNewsId(newsId).Return(0, fmt.Errorf("%s", "getAllQueryError"))
		newsInteractorMock := interactor.NewsInteractor{mockObj}
		result, errResult := newsInteractorMock.Delete(newsId)
		if result != 0 || errResult == nil {
			t.Errorf("Cannot get error from Delete news")
		}
	}()
}

func TestAddNewNews(t *testing.T) {
	mockObj := SetupNewsTest(t)
	newsItem := domain.NewsPost{
		Title:       "test",
		Description: "This is test news",
		PhotoLink:   "test@test.com",
	}

	func() {
		success := true
		mockObj.EXPECT().AddNewsItem(newsItem).Return(success, nil)
		newsInteractorMock := interactor.NewsInteractor{mockObj}
		result, errResult := newsInteractorMock.AddNewNews(newsItem)
		if result != success || errResult != nil {
			t.Errorf("Cannot add news successfully")
		}
	}()

	func() {
		success := false
		mockObj.EXPECT().AddNewsItem(newsItem).Return(success, nil)
		newsInteractorMock := interactor.NewsInteractor{mockObj}
		result, errResult := newsInteractorMock.AddNewNews(newsItem)
		if result != success || errResult != nil {
			t.Errorf("Cannot fail add news")
		}
	}()

	func() {
		success := false
		mockObj.EXPECT().AddNewsItem(newsItem).Return(success, fmt.Errorf("%s", "addNewsItemQueryError"))
		newsInteractorMock := interactor.NewsInteractor{mockObj}
		result, errResult := newsInteractorMock.AddNewNews(newsItem)
		if result != success || errResult == nil {
			t.Errorf("Cannot get error from addNewNews")
		}
	}()
}
