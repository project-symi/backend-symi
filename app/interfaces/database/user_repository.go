package database

import (
	"project-symi-backend/app/domain"
	"time"
)

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) FindKeyIdByEmployeeId(employeeId string) (id int, err error) {
	row, err := repo.Query(`SELECT id FROM users WHERE employee_id = ?`, employeeId)
	defer row.Close()
	if err != nil {
		return
	}
	row.Next()
	if err = row.Scan(&id); err != nil {
		return
	}
	return
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
			p.name
  		FROM users u
  		JOIN permissions p ON p.id = u.permission_id
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
			employeeId  string
			mail        string
			department  string
			name        string
			dateOfBirth string
			gender      string
			permission  string
		)
		if err := rows.Scan(
			&employeeId,
			&mail,
			&department,
			&name,
			&dateOfBirth,
			&gender,
			&permission); err != nil {
			continue
		}
		user := domain.User{
			EmployeeId:  employeeId,
			Name:        name,
			Mail:        mail,
			Department:  department,
			DateOfBirth: dateOfBirth,
			Gender:      gender,
			Permission:  permission,
		}
		users = append(users, user)
	}
	return
}

func (repo *UserRepository) FindTopPointsUsers(limit int) (users domain.UsersWithPoints, err error) {
	rows, err := repo.Query(`
		SELECT
			u.id,
			u.name,
			u.total_points,
			d.name,
			g.gender
  		FROM users u
  		JOIN departments d ON d.id = u.department_id
  		JOIN genders g ON g.id = u.gender_id
  		WHERE
			u.deleted = false
		ORDER BY u.total_points DESC
		LIMIT ?
		`, limit)
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var (
			id         int
			name       string
			points     int
			department string
			gender     string
		)
		if err := rows.Scan(
			&id,
			&name,
			&points,
			&department,
			&gender); err != nil {
			continue
		}
		user := domain.UserWithPoints{
			Id:         id,
			Name:       name,
			Points:     points,
			Department: department,
			Gender:     gender,
		}
		users = append(users, user)
	}
	return
}

func (repo *UserRepository) FindByEmployeeId(id string) (user domain.User, err error) {
	row, err := repo.Query(`
		SELECT
			u.employee_id,
			u.mail,
			d.name,
			u.name,
			u.birthday,
			g.gender,
			p.name
  		FROM users u
  		JOIN permissions p ON p.id = u.permission_id
  		JOIN departments d ON d.id = u.department_id
  		JOIN genders g ON g.id = u.gender_id
  		WHERE
			u.deleted = false
		AND u.employee_id = ?
		`, id)
	defer row.Close()
	if err != nil {
		return
	}
	row.Next()
	var (
		employeeId  string
		mail        string
		department  string
		name        string
		dateOfBirth string
		gender      string
		permission  string
	)
	if err = row.Scan(
		&employeeId,
		&mail,
		&department,
		&name,
		&dateOfBirth,
		&gender,
		&permission); err != nil {
		return
	}
	user = domain.User{
		EmployeeId:  employeeId,
		Name:        name,
		Mail:        mail,
		Department:  department,
		DateOfBirth: dateOfBirth,
		Gender:      gender,
		Permission:  permission,
	}
	return
}

func (repo *UserRepository) FilterByName(query string) (users domain.Users, err error) {
	rows, err := repo.Query(query)
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var (
			employeeId string
			name       string
			department string
		)
		if err := rows.Scan(
			&employeeId,
			&name,
			&department,
		); err != nil {
			continue
		}
		user := domain.User{
			EmployeeId: employeeId,
			Name:       name,
			Department: department,
		}
		users = append(users, user)
	}
	return
}

func (repo *UserRepository) DeleteByEmployeeId(id string) (amountOfDeleted int, err error) {
	result, err := repo.Execute(`
		UPDATE users
		SET deleted = true,
			deleted_at = ?
		WHERE employee_id = ?
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

func (repo *UserRepository) IsUser(employee_id string) (isUser bool, err error) {
	row, err := repo.Query("SELECT id FROM users WHERE employee_id = ?", employee_id)
	if err != nil {
		return
	}
	isUser = row.Next()
	return
}

func (repo *UserRepository) ExecuteUsersQuery(query string) (amountOfAffected int, err error) {
	result, err := repo.Execute(query)
	if err != nil {
		return
	}
	amountOfStored64, err := result.RowsAffected()
	if err != nil {
		return
	}
	amountOfAffected = int(amountOfStored64)
	return
}

func (repo *UserRepository) AddUser(employee_id string, name string, mail string, birthday string, gender_id int, department_id int, permission_id int, passwordHash string) (success bool, err error) {
	result, err := repo.Execute(`
	INSERT INTO
		users
	(employee_id, name, mail, birthday, gender_id, department_id, permission_id, created_at, modified_at, password)
	VALUES
	(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		employee_id, name, mail, birthday, gender_id, department_id, permission_id, time.Now(), time.Now(), passwordHash)
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
