package usecase

import (
	"project-symi-backend/app/domain"
)

type NewsRepository interface {
	AddNewsItem(string, string, string) (bool, error)
	GetAll() (domain.News, error)
	DeleteByNewsId(string) (int, error)
}
