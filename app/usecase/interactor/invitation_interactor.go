package interactor

import (
	"project-symi-backend/app/domain"
	"project-symi-backend/app/usecase/repository"
)

type InvitationInteractor struct {
	InvitationRepository repository.InvitationRepository
}

func (interactor *InvitationInteractor) FindBySenderId(senderId string) (invitations domain.Invitations, err error) {
	invitations, err = interactor.InvitationRepository.FindBySenderId(senderId)
	if err != nil {
		return
	}
	return
}

func (interactor *InvitationInteractor) FindById(invitationId int) (invitation domain.Invitation, err error) {
	invitation, err = interactor.InvitationRepository.FindById(invitationId)
	if err != nil {
		return
	}
	return
}
