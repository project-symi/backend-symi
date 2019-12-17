package database

import (
	"project-symi-backend/app/domain"
	"time"
)

type InvitationRepository struct {
	SqlHandler
}

func (repo *InvitationRepository) FindAll() (invitations domain.Invitations, err error) {
	rows, err := repo.Query(`
		SELECT
			i.id,
			u2.employee_id,
			i.comments,
			ic.status,
			i.reply,
			i.seen,
			i.invitation_date
		FROM invitations i
		JOIN users u1 ON u1.id = i.sender_id
		JOIN users u2 ON u2.id = i.employee_id
		JOIN invitation_status_categories ic ON ic.id = i.invitation_status_category_id
		WHERE i.deleted = false
		AND i.invitation_date >= ?`,
		time.Now().Format("2006-01-02"))
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var (
			id                 int
			employeeId         string
			comments           string
			status             string
			reply              string
			seen               bool
			invitationDateTime string
		)
		if err := rows.Scan(
			&id,
			&employeeId,
			&comments,
			&status,
			&reply,
			&seen,
			&invitationDateTime); err != nil {
			continue
		}
		dateTime, _ := time.Parse("2006-01-02 15:04:05", invitationDateTime)
		invitation := domain.Invitation{
			Id:             id,
			EmployeeId:     employeeId,
			Comments:       comments,
			Status:         status,
			Reply:          reply,
			Seen:           seen,
			InvitationDate: dateTime.Format("2006-01-02"),
			InvitationTime: dateTime.Format("15:04"),
		}
		invitations = append(invitations, invitation)
	}
	return
}

func (repo *InvitationRepository) UpdateSeenFromStatus(pending int) (err error) {
	_, err = repo.Execute(`UPDATE invitations SET seen = true WHERE deleted = false AND seen = false AND invitation_status_category_id != ?`, pending)
	return
}

func (repo *InvitationRepository) FindByEmployeeId(employeeKeyId int) (invitations domain.Invitations, err error) {
	rows, err := repo.Query(`
		SELECT
			i.id,
			u.employee_id,
			i.comments,
			ic.status,
			i.reply,
			i.seen,
			i.invitation_date
		FROM invitations i
		JOIN users u ON u.id = i.employee_id
		JOIN invitation_status_categories ic ON ic.id = i.invitation_status_category_id
		WHERE i.deleted = false
		AND i.employee_id = ?`,
		employeeKeyId)
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var (
			id                 int
			employeeId         string
			comments           string
			status             string
			reply              string
			seen               bool
			invitationDateTime string
		)
		if err = rows.Scan(
			&id,
			&employeeId,
			&comments,
			&status,
			&reply,
			&seen,
			&invitationDateTime); err != nil {
			return
		}
		dateTime, _ := time.Parse("2006-01-02 15:04:05", invitationDateTime)
		invitation := domain.Invitation{
			Id:             id,
			EmployeeId:     employeeId,
			Comments:       comments,
			Status:         status,
			Reply:          reply,
			Seen:           seen,
			InvitationDate: dateTime.Format("2006-01-02"),
			InvitationTime: dateTime.Format("15:04"),
		}
		invitations = append(invitations, invitation)
	}
	return
}
