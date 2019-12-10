package database

import (
	"errors"
	"project-symi-backend/app/domain"
	"time"

	uuid "github.com/google/uuid"
)

// uuid "github.com/satori/go.uuid"

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
  		from users u
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
  		from users u
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

//*****************************************************//
//*****IMPLEMENTING THE AUTHENTIFICATION FEATURES!*****//
//*****************************************************//

func (repo *UserRepository) IssueToken(employeeId string, employeePass string) (tokenId uuid.UUID, err error) {
	//GENERATE TOKEN
	tokenId, err = uuid.NewRandom()

	//CHECK LOGIN INFO
	row, err := repo.Query(`
	SELECT
	u.password
	from users u
	WHERE
	u.employee_id = ?
	`, employeeId)
	defer row.Close()

	if err != nil {
		return
	}

	row.Next()
	var (
		pass string
	)
	if err = row.Scan(
		&pass); err != nil {
		err = errors.New("Username Not Found")
		return
	}

	if pass != employeePass {
		err = errors.New("Password Not Found")
		return
	}
	return
}

func (repo *UserRepository) RegisterToken(employeeId string, tokenId uuid.UUID) (amountOfAffected int, err error) {
	result, err := repo.Execute(`
		UPDATE users
		SET current_token = ?
		WHERE employee_id = ?
		`, tokenId, employeeId)
	if err != nil {
		return
	}
	amountOfUpdated64, err := result.RowsAffected()
	if err != nil {
		return
	}
	amountOfAffected = int(amountOfUpdated64)
	return
}

func (repo *UserRepository) ValidateToken(tokenId uuid.UUID) (isValid bool) {
	//CHECK TOKEN ID INFO
	row, err := repo.Query(`
	SELECT
	u.current_token
	from users u
	WHERE
	u.current_token = ?
	`, tokenId)
	defer row.Close()

	if err != nil {
		return false
	}

	row.Next()
	var (
		tokenString string
	)
	if err = row.Scan(
		&tokenString); err != nil {
		err = errors.New("Error in DB")
		return false
	}

	if tokenString != tokenId.String() {
		err = errors.New("Invalid Session ID or Permission Level")
		return false
	}
	return true
}

func (repo *UserRepository) RevokeToken(tokenId uuid.UUID) (amountOfDeleted int, err error) {
	result, err := repo.Execute(`
		UPDATE users
		SET current_token = null
		WHERE current_token = ?
		AND deleted = false
		`, tokenId)
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
