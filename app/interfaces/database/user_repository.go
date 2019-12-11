package database

//TODO: Add custom errors in way that wouldnt require importing "errors" module here (used on lines: 253, 257, 326, 322 )
import (
	"errors"
	"project-symi-backend/app/domain"
	"time"

	uuid "github.com/google/uuid"
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

func (repo *UserRepository) FindTopPointUsers(limit int) (users domain.UsersWithPoint, err error) {
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
			point      int
			department string
			gender     string
		)
		if err := rows.Scan(
			&id,
			&name,
			&point,
			&department,
			&gender); err != nil {
			continue
		}
		user := domain.UserWithPoint{
			Id:         id,
			Name:       name,
			Point:      point,
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

func (repo *UserRepository) AddUser(employee_id string, name string, mail string, birthday string, gender_id int, department_id int, permission_id int) (success bool, err error) {
	result, err := repo.Execute(`
	INSERT INTO
		users
	(employee_id, name, mail, birthday, gender_id, department_id, permission_id, created_at, modified_at)
	VALUES
	(?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		employee_id, name, mail, birthday, gender_id, department_id, permission_id, time.Now(), time.Now())
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

//*****************************************************//
//*****IMPLEMENTING THE AUTHENTIFICATION FEATURES!*****//
//*****************************************************//

func (repo *UserRepository) IssueToken(employeeId string, employeePass string) (tokenId uuid.UUID, err error) {

	//CHECK LOGIN INFO
	row, err := repo.Query(`
	SELECT
	u.password
	FROM users u
	WHERE
	u.employee_id = ?
	`, employeeId)
	defer row.Close()

	if err != nil {
		return
	}

	row.Next()
	var pass string
	if err = row.Scan(&pass); err != nil {
		err = errors.New("Username Not Found")
		return
	}
	if pass != employeePass {
		err = errors.New("Incorrect Password")
		return
	}
	//GENERATE TOKEN
	tokenId, err = uuid.NewRandom()
	return tokenId, nil
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

func (repo *UserRepository) GetPermissionName(employeeId string) (permissionLevel string, err error) {
	row, err := repo.Query(`
	SELECT
	p.name
	FROM permissions p
	JOIN users u ON p.id = u.permission_id
	WHERE
	u.employee_id = ?
	`, employeeId)
	defer row.Close()

	if err != nil {
		return
	}

	row.Next()
	var permission string
	if err = row.Scan(&permission); err != nil {
		return
	}
	permissionLevel = permission
	return

}

func (repo *UserRepository) ValidateToken(tokenId uuid.UUID) (isValid bool, err error) {
	//CHECK TOKEN ID INFO
	row, err := repo.Query(`
	SELECT
	u.current_token
	FROM users u
	WHERE
	u.current_token = ?
	`, tokenId)
	defer row.Close()

	if err != nil {
		return
	}

	row.Next()
	var tokenString string
	if err = row.Scan(
		&tokenString); err != nil {
		err = errors.New("Error in DB")
		return
	}

	isValid = tokenString == tokenId.String()
	if !isValid {
		err = errors.New("Invalid Session ID or Permission Level")
		return
	}
	return
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
