package domain

type Invitation struct {
	Id             int    `json:"invitationId"`
	EmployeeId     string `json:"employeeId"`
	Comments       string `json:"comments"`
	Status         string `json:"status"`
	Reply          string `json:"reply"`
	Seen           bool   `json:"seen"`
	InvitationDate string `json:"invitationDate"`
	InvitationTime string `json:"invitationTime"`
}

type Invitations []Invitation

type PostInvitation struct {
	EmployeeId     string `json:"employeeId"`
	Comments       string `json:"comments"`
	InvitationDate string `json:"invitationDate"`
	InvitationTime string `json:"invitationTime"`
}

type PatchInvitation struct {
	Seen  bool   `json:"seen"`
	Reply string `json:"reply"`
}
