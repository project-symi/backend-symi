package database

import (
	"project-symi-backend/app/domain"
	"time"
)

type NewsRepository struct {
	SqlHandler
}

func (repo *NewsRepository) GetAllNews() (news domain.News, err error) {
	rows, err := repo.Query(`
		SELECT
			n.id,
			n.title,
			n.description,
			n.photo_link,
			n.created_at,
			g.modified_at,
  		FROM news n
  		WHERE
			u.deleted = false
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
			photo_link  string
			createdAt   string
			modifiedAt  string
		)
		if err := rows.Scan(
			&newsId,
			&title,
			&description,
			&photo_link,
			&createdAt,
			&modifiedAt); err != nil {
			continue
		}
		newsItem := domain.NewsItem{
			NewsItemId:  newsId,
			Title:       title,
			Description: description,
			PhotoLink:   photo_link,
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

func (repo *NewsRepository) AddNewsItem(title string, description string, photoLink string) (success bool, err error) {
	result, err := repo.Execute(`
	INSERT INTO
		news
	(title, description, photo_link, created_at, modified_at)
	VALUES
	(?, ?, ?, ?, ?)`,
		title, description, photoLink, time.Now(), time.Now())
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
