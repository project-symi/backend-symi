package usecase

import (
	"project-symi-backend/app/domain"

	uuid "github.com/google/uuid"
)

type UserRepository interface {
	FindAll() (domain.Users, error)
	FindByEmployeeId(employeeId string) (domain.User, error)
	FilterByName(query string) (domain.Users, error)
	DeleteByEmployeeId(employeeId string) (amountOfDeleted int, err error)
	IsUser(employeeId string) (isUser bool, err error)
	ExecuteUsersQuery(query string) (amountOfAffected int, err error)
	IssueToken(employeeId string, employeePass string) (tokenId uuid.UUID, err error)
	RegisterToken(employeeId string, tokenId uuid.UUID) (amountOfAffected int, err error)
	ValidateToken(tokenId uuid.UUID) (isValid bool)
	RevokeToken(tokenId uuid.UUID) (amountOfAffected int, err error)
}
