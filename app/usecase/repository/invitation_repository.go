package repository

import (
	"project-symi-backend/app/domain"
)

type InvitationRepository interface {
	UpdateSeenFromStatus(int) error
	FindAll() (domain.LeaderInvitations, error)
	FindByEmployeeId(int) (domain.EmployeeInvitations, error)
	UpdateStatusAndReplyById(int, int, string) (bool, error)
	PostInvitation(int, int, string, string) (bool, error)
	DeleteById(int) (bool, error)
}
