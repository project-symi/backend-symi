package usecase

import (
	"project-symi-backend/app/domain"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserInteractor struct {
	UserRepository       UserRepository
	GenderRepository     GenderRepository
	DepartmentRepository DepartmentRepository
	PermissionRepository PermissionRepository
}

func (interactor *UserInteractor) Users() (user domain.Users, err error) {
	user, err = interactor.UserRepository.FindAll()
	return
}

func (interactor *UserInteractor) FindTopPointsUsers(numOfRank int) (users domain.UsersWithPoints, err error) {
	users, err = interactor.UserRepository.FindTopPointsUsers(numOfRank)
	return
}

func (interactor *UserInteractor) User(employeeId string) (user domain.User, err error) {
	user, err = interactor.UserRepository.FindByEmployeeId(employeeId)
	return
}

func (interactor *UserInteractor) UsersByName(name string) (users domain.Users, err error) {
	nameArray := strings.Split(name, " ")
	query := createFilterByNameQuery(nameArray)
	users, err = interactor.UserRepository.FilterByName(query)
	return
}

func createFilterByNameQuery(nameArray []string) (query string) {
	query = `
		SELECT
			u.employee_id,
			u.name,
			d.name
		FROM users u
		JOIN departments d ON d.id = u.department_id
		WHERE
			u.deleted = false
		`
	for _, s := range nameArray {
		query += " AND u.name LIKE '%" + s + "%'"
	}
	return
}

func (interactor *UserInteractor) Delete(employeeId string) (amountOfDeleted int, err error) {
	amountOfDeleted, err = interactor.UserRepository.DeleteByEmployeeId(employeeId)
	return
}

func (interactor *UserInteractor) StoreUser(user domain.User) (success bool, err error) {
	genders, err := interactor.GenderRepository.FindAll()
	departments, err := interactor.DepartmentRepository.FindAll()
	permissions, err := interactor.PermissionRepository.FindAll()
	if err != nil {
		return
	}
	success, err = interactor.UserRepository.AddUser(user.EmployeeId, user.Name, user.Mail, user.DateOfBirth, genders.GenderToId(user.Gender), departments.DepartmentToId(user.Department), permissions.PermissionToId(user.Permission), passwordHash(user.DateOfBirth))
	return
}

func (interactor *UserInteractor) StoreUsers(users domain.Users) (amountOfChanged int, err error) {
	var (
		amountOfInserted int = 0
		amountOfUpdated  int = 0
	)
	registeredUsers, unRegisteredUsers, err := interactor.divideRegisteredAndUnregisteredUsers(users)
	if len(registeredUsers) == 0 && len(unRegisteredUsers) == 0 {
		amountOfChanged = 0
		return
	}
	insertQuery, updateQuery, err := interactor.createStoreUsersQuery(registeredUsers, unRegisteredUsers)
	if len(unRegisteredUsers) != 0 {
		amountOfInserted, err = interactor.UserRepository.ExecuteUsersQuery(insertQuery)
	}
	if len(registeredUsers) != 0 {
		amountOfUpdated, err = interactor.UserRepository.ExecuteUsersQuery(updateQuery)
	}
	amountOfChanged = amountOfInserted + amountOfUpdated
	return
}

func (interactor *UserInteractor) divideRegisteredAndUnregisteredUsers(users domain.Users) (registeredUsers domain.Users, unRegisteredUsers domain.Users, err error) {
	isUser := false
	for _, user := range users {
		isUser, err = interactor.UserRepository.IsUser(user.EmployeeId)
		if isUser {
			registeredUsers = append(registeredUsers, user)
		} else {
			unRegisteredUsers = append(unRegisteredUsers, user)
		}
	}
	return
}

func (interactor *UserInteractor) createStoreUsersQuery(registeredUsers domain.Users, unregisteredUsers domain.Users) (insertQuery string, updateQuery string, err error) {
	genders, err := interactor.GenderRepository.FindAll()
	departments, err := interactor.DepartmentRepository.FindAll()
	permissions, err := interactor.PermissionRepository.FindAll()
	if err != nil {
		return
	}
	insertQuery = createInsertQuery(unregisteredUsers, genders, departments, permissions)
	updateQuery = createUpdateQuery(registeredUsers, genders, departments, permissions)
	return
}

func createInsertQuery(users domain.Users, genders domain.Genders, departments domain.Departments, permissions domain.Permissions) (query string) {
	query = "INSERT INTO users (employee_id, name, mail, birthday, gender_id, department_id, permission_id, deleted, created_at, modified_at, password) VALUES "
	for i, user := range users {
		gender_id := genders.GenderToId(user.Gender)
		department_id := departments.DepartmentToId(user.Department)
		permission_id := permissions.PermissionToId(user.Permission)
		query += "( \"" + user.EmployeeId + "\", \"" + user.Name + "\", \"" + user.Mail + "\", \"" + user.DateOfBirth + "\", \"" + strconv.Itoa(gender_id) + "\", \"" + strconv.Itoa(department_id) + "\", \"" + strconv.Itoa(permission_id) + "\", " + "false" + ", \"" + time.Now().Format("2006-01-02 15:04:05") + "\", \"" + time.Now().Format("2006-01-02 15:04:05") + "\", \"" + passwordHash(user.DateOfBirth) + "\")"
		if i != len(users)-1 {
			query += ", "
		}
	}
	return
}

func createUpdateQuery(users domain.Users, genders domain.Genders, departments domain.Departments, permissions domain.Permissions) (query string) {
	query = "UPDATE users SET "
	nameQuery := "name = CASE employee_id "
	mailQuery := "mail = CASE employee_id "
	birthdayQuery := "birthday = CASE employee_id "
	genderQuery := "gender_id = CASE employee_id "
	departmentQuery := "department_id = CASE employee_id "
	permissionQuery := "permission_id = CASE employee_id "
	modifiedQuery := "modified_at = CASE employee_id "
	whereQuery := "WHERE employee_id IN ("
	for i, user := range users {
		gender_id := genders.GenderToId(user.Gender)
		department_id := departments.DepartmentToId(user.Department)
		permission_id := permissions.PermissionToId(user.Permission)
		nameQuery += "WHEN \"" + user.EmployeeId + "\" THEN \"" + user.Name + "\" "
		mailQuery += "WHEN \"" + user.EmployeeId + "\" THEN \"" + user.Mail + "\" "
		birthdayQuery += "WHEN \"" + user.EmployeeId + "\" THEN \"" + user.DateOfBirth + "\" "
		genderQuery += "WHEN \"" + user.EmployeeId + "\" THEN \"" + strconv.Itoa(gender_id) + "\" "
		departmentQuery += "WHEN \"" + user.EmployeeId + "\" THEN \"" + strconv.Itoa(department_id) + "\" "
		permissionQuery += "WHEN \"" + user.EmployeeId + "\" THEN \"" + strconv.Itoa(permission_id) + "\" "
		modifiedQuery += "WHEN \"" + user.EmployeeId + "\" THEN \"" + time.Now().Format("2006-01-02 15:04:05") + "\" "
		if i != len(users)-1 {
			whereQuery += "\"" + user.EmployeeId + "\", "
		}
		if i == len(users)-1 {
			nameQuery += "END, "
			mailQuery += "END, "
			birthdayQuery += "END, "
			genderQuery += "END, "
			departmentQuery += "END, "
			permissionQuery += "END, "
			modifiedQuery += "END "
			whereQuery += "\"" + user.EmployeeId + "\")"
		}
	}
	query += nameQuery + mailQuery + birthdayQuery + genderQuery + departmentQuery + permissionQuery + modifiedQuery + whereQuery
	return
}

//UTIL FUNCTIONS//

func passwordHash(pass string) (hashedPassword string) {
	passwordToHash := []byte(pass)

	// Hashing the password with the default cost of 10
	hash, err := bcrypt.GenerateFromPassword(passwordToHash, bcrypt.DefaultCost)
	if err != nil {
		return
	}
	hashedPassword = string(hash)
	return
}
