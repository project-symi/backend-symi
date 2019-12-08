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

type Users []User
