package repository

import "project-symi-backend/app/domain"

type SlackRepository interface {
	FindByName(string) (domain.Slack, error)
}
