package interactor

import (
	"project-symi-backend/app/domain"
	"project-symi-backend/app/usecase/repository"
)

const Pending = 1

type InvitationInteractor struct {
	InvitationRepository               repository.InvitationRepository
	UserRepository                     repository.UserRepository
	InvitationStatusCategoryRepository repository.InvitationStatusCategoryRepository
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

func (interactor *InvitationInteractor) PostInvitation(invitation domain.PostInvitation) (invitations domain.Invitations, err error) {
	employeeId, err := interactor.UserRepository.FindKeyIdByEmployeeId(invitation.EmployeeId)
	ceoId, err := interactor.UserRepository.FindCEOId()
	invitationDate := invitation.InvitationDate + " " + invitation.InvitationTime
	success, err := interactor.InvitationRepository.PostInvitation(ceoId, employeeId, invitation.Comments, invitationDate)
	if success == true {
		invitations, err = interactor.InvitationRepository.FindAll()
	}
	if err != nil {
		return
	}
	return
}

func (interactor *InvitationInteractor) PatchInvitationById(id int, statusAndReply domain.PatchInvitation) (success bool, err error) {
	statusId, err := interactor.InvitationStatusCategoryRepository.FindKeyIdByStatus(statusAndReply.Status)
	success, err = interactor.InvitationRepository.UpdateStatusAndReplyById(id, statusId, statusAndReply.Reply)
	if err != nil {
		return
	}
	return
}
