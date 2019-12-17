package interactor

import (
	"project-symi-backend/app/domain"
	"project-symi-backend/app/usecase/repository"
)

const Pending = 1

type InvitationInteractor struct {
	InvitationRepository repository.InvitationRepository
	UserRepository       repository.UserRepository
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

func (interactor *InvitationInteractor) FindByEmployeeId(employeeId string) (invitations domain.Invitations, err error) {
	keyId, err := interactor.UserRepository.FindKeyIdByEmployeeId(employeeId)
	invitations, err = interactor.InvitationRepository.FindByEmployeeId(keyId)
	if err != nil {
		return
	}
	return
}
