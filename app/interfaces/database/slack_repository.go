package database

import (
	"project-symi-backend/app/domain"
)

type SlackRepository struct {
	SqlHandler
}

func (repo *SlackRepository) FindByName(name string) (slack domain.Slack, err error) {
	row, err := repo.Query("SELECT token, url, text FROM slacks WHERE deleted = false AND name = ?", name)
	defer row.Close()
	if err != nil {
		return
	}
	row.Next()
	var (
		token string
		url   string
		text  string
	)
	if err = row.Scan(
		&token,
		&url,
		&text); err != nil {
		return
	}
	slack = domain.Slack{
		Token: token,
		Url:   url,
		Text:  text,
	}
	return
}
