package interactor

import (
	"project-symi-backend/app/domain"
	"project-symi-backend/app/usecase/repository"
)

const Pending = 1

type InvitationInteractor struct {
	InvitationRepository repository.InvitationRepository
}

func (interactor *InvitationInteractor) ChangeSeenAndFindAll() (invitations domain.Invitations, err error) {
	err = interactor.InvitationRepository.UpdateSeenFromStatus(Pending)
	if err != nil {
		return
	}
	invitations, err = interactor.InvitationRepository.FindAll()
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
