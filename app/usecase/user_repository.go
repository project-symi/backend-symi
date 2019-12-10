package usecase

import (
	"project-symi-backend/app/domain"

	uuid "github.com/google/uuid"
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
	IssueToken(string, string) (uuid.UUID, error)
	RegisterToken(string, uuid.UUID) (int, error)
	GetPermissionName(string) (string, error)
	ValidateToken(uuid.UUID) (bool, error)
	RevokeToken(uuid.UUID) (int, error)
}
