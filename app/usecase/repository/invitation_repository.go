package repository

import (
	"project-symi-backend/app/domain"
)

type InvitationRepository interface {
	FindBySenderId(string) (domain.Invitations, error)
	FindById(int) (domain.Invitation, error)
}
