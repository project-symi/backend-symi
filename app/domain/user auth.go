package domain

type UserCredentials struct {
	EmployeeId string `json:"userId"`
	Pass       string `json:"password"`
}

type TokenResponse struct {
	Token           string `json:"token"`
	PermissionLevel string `json:"permission"`
}
