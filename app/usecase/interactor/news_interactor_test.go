package interactor_test

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"project-symi-backend/app/domain"
	"project-symi-backend/app/usecase/interactor"
	"project-symi-backend/app/usecase/mock_repository"
	"testing"
)

func TestNews(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockObj := mock_repository.NewMockNewsRepository(mockCtrl)
	defer mockCtrl.Finish()

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
