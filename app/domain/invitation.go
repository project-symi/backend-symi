package domain

type LeaderInvitation struct {
	Id             int    `json:"invitationId"`
	EmployeeId     string `json:"employeeId"`
	EmployeeName   string `json:"employeeName"`
	Comments       string `json:"comments"`
	Status         string `json:"status"`
	Reply          string `json:"reply"`
	Seen           bool   `json:"seen"`
	InvitationDate string `json:"invitationDate"`
	InvitationTime string `json:"invitationTime"`
}

type LeaderInvitations []LeaderInvitation

type EmployeeInvitation struct {
	Id             int    `json:"invitationId"`
	Comments       string `json:"comments"`
	Status         string `json:"status"`
	Reply          string `json:"reply"`
	Seen           bool   `json:"seen"`
	InvitationDate string `json:"invitationDate"`
	InvitationTime string `json:"invitationTime"`
}

type EmployeeInvitations []EmployeeInvitation

type PostInvitation struct {
	EmployeeId     string `json:"employeeId"`
	Comments       string `json:"comments"`
	InvitationDate string `json:"invitationDate"`
	InvitationTime string `json:"invitationTime"`
}

type PatchInvitation struct {
	Status string `json:"status"`
	Reply  string `json:"reply"`
}
