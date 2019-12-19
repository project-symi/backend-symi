package repository

import (
	"project-symi-backend/app/domain"
)

type UserRepository interface {
	FindKeyIdByEmployeeId(string) (int, error)
	FindKeyIdAndSlackIdByEmployeeId(string) (int, string, error)
	FindCEOId() (int, error)
	FindAll() (domain.Users, error)
	FindTopPointsUsers(int) (domain.UsersWithPoints, error)
	FindByEmployeeId(string) (domain.UserInfoWithPoints, error)
	FilterByName(string) (domain.UsersByName, error)
	DeleteByEmployeeId(string) (int, error)
	IsUser(string) (bool, error)
	AddUser(string, string, string, string, int, int, int, string) (bool, error)
	ExecuteUsersQuery(string) (int, error)
}
