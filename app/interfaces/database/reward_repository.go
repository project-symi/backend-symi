package database

import (
	"project-symi-backend/app/domain"
)

type RewardRepository struct {
	SqlHandler
}

func (repo *RewardRepository) FindAll() (rewards domain.Rewards, err error) {
	rows, err := repo.Query("SELECT id, name, points, url FROM rewards WHERE deleted = false")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var (
			id     int
			name   string
			points int
			url    string
		)
		if err := rows.Scan(
			&id,
			&name,
			&points,
			&url); err != nil {
			continue
		}
		reward := domain.Reward{
			Id:     id,
			Name:   name,
			Points: points,
			Url:    url,
		}
		rewards = append(rewards, reward)
	}
	return
}

func (repo *RewardRepository) PatchById(reward domain.Reward) (success bool, err error) {
	result, err := repo.Execute("UPDATE rewards SET name = ?, points = ?, url = ? WHERE id = ?", reward.Name, reward.Points, reward.Url, reward.Id)
	if err != nil {
		return
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		success = false
		return
	}
	success = true
	return
}
