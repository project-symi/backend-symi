package usecase

import (
	"project-symi-backend/app/domain"
	"strconv"
	"strings"
	"time"
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

func (interactor *UserInteractor) User(employeeId string) (user domain.User, err error) {
	user, err = interactor.UserRepository.FindByEmployeeId(employeeId)
	return
}

func (interactor *UserInteractor) UsersByName(name string) (users domain.Users, err error) {
	nameArray := strings.Split(name, " ")
	users, err = interactor.UserRepository.FilterByName(nameArray)
	return
}

func (interactor *UserInteractor) Delete(employeeId string) (amountOfDeleted int, err error) {
	amountOfDeleted, err = interactor.UserRepository.DeleteByEmployeeId(employeeId)
	return
}

func (interactor *UserInteractor) Store(users domain.Users) (amountOfChanged int, err error) {
	var (
		amountOfInserted int = 0
		amountOfUpdated  int = 0
		storeQuery       string
	)
	registeredUsers, unRegisteredUsers, err := interactor.divideRegisteredAndUnregisteredUsers(users)
	if len(registeredUsers) == 0 && len(unRegisteredUsers) == 0 {
		amountOfChanged = 0
		return
	}
	if len(unRegisteredUsers) != 0 {
		storeQuery, err = interactor.createStoreUsersQuery(unRegisteredUsers)
		amountOfInserted, err = interactor.UserRepository.StoreUsers(storeQuery)
	}
	// if len(registeredUsers) != 0 {
	// 	updateQuery, err = interactor.createUpdateUsersQuery(registeredUsers)
	// 	amountOfUpdated, err = interactor.UserRepository.BulkUpdateUsers(updateQuery)
	// }
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

func (interactor *UserInteractor) createStoreUsersQuery(users domain.Users) (query string, err error) {
	genders, err := interactor.GenderRepository.FindAll()
	departments, err := interactor.DepartmentRepository.FindAll()
	permissions, err := interactor.PermissionRepository.FindAll()
	if err != nil {
		return
	}
	query = "INSERT INTO users (employee_id, name, mail, birthday, gender_id, department_id, permission_id, deleted, created_at, modified_at) VALUES "
	for i, user := range users {
		gender_id := genders.GenderToId(user.Gender)
		department_id := departments.DepartmentToId(user.Department)
		permission_id := permissions.PermissionToId(user.Permission)
		query += "( \"" + user.EmployeeId + "\", \"" + user.Name + "\", \"" + user.Mail + "\", \"" + user.DateOfBirth + "\", \"" + strconv.Itoa(gender_id) + "\", \"" + strconv.Itoa(department_id) + "\", \"" + strconv.Itoa(permission_id) + "\", " + "false" + ", \"" + time.Now().Format("2006-01-02 15:04:05") + "\", \"" + time.Now().Format("2006-01-02 15:04:05") + "\")"
		if i != len(users)-1 {
			query += ", "
		}
		if err != nil {
			return
		}
	}
	return
}
