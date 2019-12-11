package usecase

import (
	"project-symi-backend/app/domain"
)

type UserRepository interface {
	FindKeyIdByEmployeeId(string) (int, error)
	FindAll() (domain.Users, error)
	FindByEmployeeId(string) (domain.User, error)
	FilterByName(string) (domain.Users, error)
	DeleteByEmployeeId(string) (int, error)
	IsUser(string) (bool, error)
	AddUser(string, string, string, string, int, int, int) (bool, error)
	ExecuteUsersQuery(string) (int, error)
	IssueToken(string, string) (string, error)
	RegisterToken(string, string) (int, error)
	GetPermissionName(string) (string, error)
	ValidateToken(string) (bool, error)
	RevokeToken(string) (int, error)
}
