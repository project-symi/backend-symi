package domain

type User struct {
	EmployeeId  string `json:"employeeId"`
	Name        string `json:"name"`
	Mail        string `json:"email"`
	Department  string `json:"department"`
	DateOfBirth string `json:"dateOfBirth"`
	Gender      string `json:"gender"`
	Permission  string `json:"permission"`
}

type UserInfoWithPoints struct {
	EmployeeId  string `json:"employeeId"`
	Name        string `json:"name"`
	Mail        string `json:"email"`
	Department  string `json:"department"`
	DateOfBirth string `json:"dateOfBirth"`
	TotalPoints int    `json:"totalPoints"`
	Gender      string `json:"gender"`
	Permission  string `json:"permission"`
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

type Users []User
type UsersWithPoints []UserWithPoints
type UsersByName []UserByName
