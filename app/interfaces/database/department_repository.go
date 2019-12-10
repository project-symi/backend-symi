package database

import "project-symi-backend/app/domain"

type DepartmentRepository struct {
	SqlHandler
}

func (repo *DepartmentRepository) FindAll() (departments domain.Departments, err error) {
	rows, err := repo.Query("SELECT id, name FROM departments WHERE deleted = false")
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
		department := domain.Department{
			Id:   id,
			Name: name,
		}
		departments = append(departments, department)
	}
	return
}
