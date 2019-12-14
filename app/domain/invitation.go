package domain

type Invitation struct {
	Id             int    `json:"invitationId"`
	EmployeeId     string `json:"employeeId"`
	Comments       string `json:"comments"`
	Status         string `json:"status"`
	Reply          string `json:"reply"`
	Seen           bool   `json:"seen"`
	InvitationDate string `json:"invitationDate"`
}

type InvitationPost struct {
	SenderId       string `json:"senderId"`
	EmployeeId     string `json:"employeeId"`
	Comments       string `json:"comments"`
	InvitationDate string `json:"invitationDate"`
}

type Invitations []Invitation
