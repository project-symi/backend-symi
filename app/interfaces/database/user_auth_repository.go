package database

//TODO: Add custom errors in way that wouldnt require importing "errors" module here (used on lines: 30, 34, 107, 113 )
import (
	"errors"

	uuid "github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

//*****************************************************//
//*****IMPLEMENTING THE AUTHENTIFICATION FEATURES!*****//
//*****************************************************//

type UserAuthRepository struct {
	SqlHandler
}

func (repo *UserAuthRepository) IssueToken(employeeId string, password string) (tokenIdStr string, err error) {

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
	var pass []byte
	if err = row.Scan(&pass); err != nil {
		err = errors.New("Username Not Found")
		return
	}

	//check the password hash
	err = bcrypt.CompareHashAndPassword(pass, []byte(password))

	if err != nil {
		err = errors.New("Incorrect Password")
		return
	}
	//GENERATE TOKEN
	tokenUUID, err := uuid.NewRandom()
	if err != nil {
		return
	}
	tokenIdStr = tokenUUID.String()
	return tokenIdStr, nil
}

func (repo *UserAuthRepository) RegisterToken(employeeId string, tokenId string) (amountOfAffected int, err error) {
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

func (repo *UserAuthRepository) GetPermissionName(employeeId string) (permissionLevel string, err error) {
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

func (repo *UserAuthRepository) ValidateToken(tokenId string) (isValid bool, err error) {
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

	isValid = tokenString == tokenId
	if !isValid {
		err = errors.New("Invalid Session ID or Permission Level")
		return
	}
	return
}

func (repo *UserAuthRepository) RevokeToken(tokenId string) (amountOfDeleted int, err error) {
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
