package repository

import (
	"project-symi-backend/app/domain"
)

type InvitationRepository interface {
	UpdateSeenFromStatus(int) error
	FindAll() (domain.Invitations, error)
	FindByEmployeeId(int) (domain.Invitations, error)
	UpdateStatusAndReplyById(int, int, string) (bool, error)
}
