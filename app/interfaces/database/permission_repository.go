package database

import "project-symi-backend/app/domain"

type PermissionRepository struct {
	SqlHandler
}

func (repo *PermissionRepository) FindAll() (permissions domain.Permissions, err error) {
	rows, err := repo.Query("SELECT id, name FROM permissions WHERE deleted = false")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var (
			id   int
			name string
		)
		if err := rows.Scan(
			&id,
			&name); err != nil {
			continue
		}
		permission := domain.Permission{
			Id:   id,
			Name: name,
		}
		permissions = append(permissions, permission)
	}
	return
}
