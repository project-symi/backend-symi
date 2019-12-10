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
	ExecuteUsersQuery(string) (int, error)
	IssueToken(employeeId string, employeePass string) (tokenId uuid.UUID, err error)
	RegisterToken(employeeId string, tokenId uuid.UUID) (amountOfAffected int, err error)
	ValidateToken(tokenId uuid.UUID) (isValid bool)
	RevokeToken(tokenId uuid.UUID) (amountOfAffected int, err error)
}
