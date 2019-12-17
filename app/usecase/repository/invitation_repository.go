package repository

import (
	"project-symi-backend/app/domain"
)

type InvitationRepository interface {
	UpdateSeenFromStatus(int) error
	FindAll() (domain.Invitations, error)
	FindById(int) (domain.Invitation, error)
}
