package database

import "project-symi-backend/app/domain"

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) FindAll() (users domain.Users, err error) {
	rows, err := repo.Query(`
		SELECT
			u.employee_id,
			u.mail,
			d.name,
			u.name,
			u.birthday,
			g.gender,
			t.name
  		from users u
  		JOIN types t ON t.id = u.type_id
  		JOIN departments d ON d.id = u.department_id
  		JOIN genders g ON g.id = u.gender_id
  		WHERE
			u.deleted = false
		`)
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var (
			employeeId  int
			mail        string
			department  string
			name        string
			dateOfBirth string
			gender      string
			typeName    string
		)
		if err := rows.Scan(
			&employeeId,
			&mail,
			&department,
			&name,
			&dateOfBirth,
			&gender,
			&typeName); err != nil {
			continue
		}
		user := domain.User{
			EmployeeId:  employeeId,
			Name:        name,
			Mail:        mail,
			Department:  department,
			DateOfBirth: dateOfBirth,
			Gender:      gender,
			Type:        typeName,
		}
		users = append(users, user)
	}
	return
}

func (repo *UserRepository) FindById(id int) (user domain.User, err error) {
	rows, err := repo.Query(`
		SELECT
			u.employee_id,
			u.mail,
			d.name,
			u.name,
			u.birthday,
			g.gender,
			t.name
  		from users u
  		JOIN types t ON t.id = u.type_id
  		JOIN departments d ON d.id = u.department_id
  		JOIN genders g ON g.id = u.gender_id
  		WHERE
			u.deleted = false
		AND u.employee_id = ?
		`, id)
	defer rows.Close()
	if err != nil {
		return
	}
	rows.Next()
	var (
		employeeId  int
		mail        string
		department  string
		name        string
		dateOfBirth string
		gender      string
		typeName    string
	)
	if err = rows.Scan(
		&employeeId,
		&mail,
		&department,
		&name,
		&dateOfBirth,
		&gender,
		&typeName); err != nil {
		return
	}
	user = domain.User{
		EmployeeId:  employeeId,
		Name:        name,
		Mail:        mail,
		Department:  department,
		DateOfBirth: dateOfBirth,
		Gender:      gender,
		Type:        typeName,
	}
	return
}
