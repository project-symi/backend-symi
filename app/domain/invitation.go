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
	EmployeeId     string `json:"employeeId" binding:"required,max=20,alphanum"`
	Comments       string `json:"comments" binding:"required"`
	InvitationDate string `json:"invitationDate" binding:"required,excludesall=!()#@{}?"` //TODO: Custom validation
	InvitationTime string `json:"invitationTime" binding:"required,excludesall=!()#@{}?"` //TODO: Custom validation
}

type PatchInvitation struct {
	Status string `json:"status" binding:"required,oneof=pending accepted declined"`
	Reply  string `json:"reply" binding:"required,max=5000"`
}

type InvitationIdParam struct {
	InvitationId int `uri:"invitationId" binding:"required,numeric"`
}
