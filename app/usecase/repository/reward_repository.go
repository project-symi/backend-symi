package repository

import "project-symi-backend/app/domain"

type RewardRepository interface {
	FindAll() (domain.Rewards, error)
	PatchById(domain.Reward) (bool, error)
}
