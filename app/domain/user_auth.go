package domain

type UserCredentials struct {
	EmployeeId string `json:"userId" binding:"required,max=20,alphanum"`
	Pass       string `json:"password" binding:"required"`
}

type TokenResponse struct {
	Token           string `json:"token"`
	PermissionLevel string `json:"permission"`
}

type TokenRequest struct {
	Token string `header:"token" binding:"required,len=204"`
}
