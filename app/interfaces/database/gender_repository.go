package database

import "project-symi-backend/app/domain"

type GenderRepository struct {
	SqlHandler
}

func (repo *GenderRepository) FindAll() (genders domain.Genders, err error) {
	rows, err := repo.Query("SELECT id, gender from genders WHERE deleted = false")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var (
			id     int
			gender string
		)
		if err := rows.Scan(&id, &gender); err != nil {
			continue
		}
		genderEntity := domain.Gender{
			Id:     id,
			Gender: gender,
		}
		genders = append(genders, genderEntity)
	}
	return
}
