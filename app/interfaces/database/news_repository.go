package database

import (
	"project-symi-backend/app/domain"
	"time"
)

type NewsRepository struct {
	SqlHandler
}

func (repo *NewsRepository) GetAll() (news domain.News, err error) {
	rows, err := repo.Query(`
		SELECT
			id,
			title,
			description,
			photo_link,
			created_at,
			modified_at
  		FROM news
  		WHERE
			deleted = false
		`)
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var (
			newsId      int
			title       string
			description string
			photoLink   string
			createdAt   string
			modifiedAt  string
		)
		if err := rows.Scan(
			&newsId,
			&title,
			&description,
			&photoLink,
			&createdAt,
			&modifiedAt); err != nil {
			continue
		}
		newsItem := domain.NewsItem{
			NewsItemId:  newsId,
			Title:       title,
			Description: description,
			PhotoLink:   photoLink,
			CreatedAt:   createdAt,
			ModifiedAt:  modifiedAt,
		}
		news = append(news, newsItem)
	}
	return
}

func (repo *NewsRepository) DeleteByNewsId(id string) (amountOfDeleted int, err error) {
	result, err := repo.Execute(`
		UPDATE news
		SET deleted = true,
			deleted_at = ?
		WHERE id = ?
		AND deleted = false
		`, time.Now(), id)
	if err != nil {
		return
	}
	amountOfDeleted64, err := result.RowsAffected()
	if err != nil {
		return
	}
	amountOfDeleted = int(amountOfDeleted64)
	return
}

func (repo *NewsRepository) AddNewsItem(news domain.NewsPost) (success bool, err error) {
	result, err := repo.Execute(`
	INSERT INTO
		news
	(title, description, photo_link, created_at, modified_at)
	VALUES
	(?, ?, ?, ?, ?)`,
		news.Title, news.Description, news.PhotoLink, time.Now(), time.Now())
	if err != nil {
		return
	}
	amountOfStored64, err := result.RowsAffected()
	if err != nil {
		return
	}
	if amountOfStored64 == 1 {
		success = true
		return
	}
	return
}
