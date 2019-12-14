package database

import (
	"project-symi-backend/app/domain"
	"time"
)

type InvitationRepository struct {
	SqlHandler
}

func (repo *InvitationRepository) FindBySenderId(employeeId string) (invitations domain.Invitations, err error) {
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
		WHERE u1.employee_id = ?
		AND i.deleted = false
		AND i.invitation_date >= ?`,
		employeeId, time.Now().Format("2006-01-02"))
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var (
			id             int
			employeeId     string
			comments       string
			status         string
			reply          string
			seen           bool
			invitationDate string
		)
		if err := rows.Scan(
			&id,
			&employeeId,
			&comments,
			&status,
			&reply,
			&seen,
			&invitationDate); err != nil {
			continue
		}
		invitation := domain.Invitation{
			Id:             id,
			EmployeeId:     employeeId,
			Comments:       comments,
			Status:         status,
			Reply:          reply,
			Seen:           seen,
			InvitationDate: invitationDate,
		}
		invitations = append(invitations, invitation)
	}
	return
}
