package domain

type User struct {
	EmployeeId  string `json:"employeeId" binding:"required,max=20,alphanum,excludesall=!()#@{}?"`
	Name        string `json:"name" binding:"required,max=20,excludesall=!()#@{}?"`
	Mail        string `json:"email" binding:"required,email"`
	Department  string `json:"department" binding:"required,excludesall=!()#@{}?"`
	DateOfBirth string `json:"dateOfBirth" binding:"required,len=10"`
	Gender      string `json:"gender" binding:"required,oneof=male female"`
	Permission  string `json:"permission" binding:"required,oneof=CEO employee admin"`
}

type Users []User

type UsersRoot struct {
	Users `json:"users" binding:"required,dive"`
}

type UserInfoWithPoints struct {
	EmployeeId  string `json:"employeeId" binding:"required,max=20,alphanum"`
	Name        string `json:"name" binding:"required,max=20,excludesall=!()#@{}?"`
	Mail        string `json:"email" binding:"required,email"`
	Department  string `json:"department" binding:"required,excludesall=!()#@{}?"`
	DateOfBirth string `json:"dateOfBirth" binding:"required,len=10,excludesall=!()#@{}?"` //TODO: Custom validation
	TotalPoints int    `json:"totalPoints" binding:"required,numeric"`
	Gender      string `json:"gender" binding:"required,oneof=male female"`
	Permission  string `json:"permission" binding:"required,oneof=CEO employee admin"`
}

type UserByName struct {
	EmployeeId string `json:"employeeId"`
	Name       string `json:"name"`
	Department string `json:"department"`
}

type UserWithPoints struct {
	EmployeeId  string `json:"employeeId"`
	Name        string `json:"name"`
	Points      int    `json:"points"`
	DateOfBirth string `json:"dateOfBirth"`
	Department  string `json:"department"`
	Gender      string `json:"gender"`
}

type UsersWithPoints []UserWithPoints
type UsersByName []UserByName
type NameQuery struct {
	Name string `form:"name" binding:"required,max=20,excludesall=!()#@{}?"`
}
type EmployeeIdParam struct {
	EmployeeId string `uri:"employeeId" binding:"required,max=20,alphanum"`
}
