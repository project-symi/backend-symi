package usecase

import "project-symi-backend/app/domain"

type UserRepository interface {
	FindAll() (domain.Users, error)
	FindByEmployeeId(employeeId string) (domain.User, error)
	FilterByName(query string) (domain.Users, error)
	DeleteByEmployeeId(employeeId string) (amountOfDeleted int, err error)
	IsUser(employeeId string) (isUser bool, err error)
	ExecuteUsersQuery(query string) (amountOfAffected int, err error)
	IssueToken(employeeId string, employeePass string) (tokenId string, err error)
	ValidateToken(tokenId string) (isValid bool)
	RevokeToken(tokenId string) (amountOfAffected int, err error)
}
