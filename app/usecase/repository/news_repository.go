package repository

import (
	"project-symi-backend/app/domain"
)

type NewsRepository interface {
	GetAll() (domain.News, error)
	DeleteByNewsId(string) (int, error)
	AddNewsItem(domain.NewsPost) (bool, error)
}
