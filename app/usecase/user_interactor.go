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

func (interactor *UserInteractor) Store(users domain.Users) (amountOfStored int, err error) {
	noDuplicateUser, err := interactor.deleteDuplicateUsers(users)
	if len(noDuplicateUser) == 0 {
		amountOfStored = 0
		return
	}
	storeQuery, err := interactor.createStoreUsersQuery(noDuplicateUser)
	amountOfStored, err = interactor.UserRepository.StoreUsers(storeQuery)
	return
}

func (interactor *UserInteractor) deleteDuplicateUsers(users domain.Users) (noDuplicateUsers domain.Users, err error) {
	isEmployee := false
	for _, user := range users {
		isEmployee, err = interactor.UserRepository.IsEmployee(user.EmployeeId)
		if !isEmployee {
			noDuplicateUsers = append(noDuplicateUsers, user)
		}
	}
	return
}

func (interactor *UserInteractor) createStoreUsersQuery(users domain.Users) (query string, err error) {
	var (
		gender_id     int
		department_id int
		permission_id int
	)
	query = "INSERT INTO users (employee_id, name, mail, birthday, gender_id, department_id, permission_id, deleted, created_at, modified_at) VALUES "
	for i, user := range users {
		gender_id, err = interactor.GenderRepository.GenderToId(user.Gender)
		department_id, err = interactor.DepartmentRepository.DepartmentToId(user.Department)
		permission_id, err = interactor.PermissionRepository.PermissionToId(user.Permission)
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
