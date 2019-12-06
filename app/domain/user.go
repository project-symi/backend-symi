package domain

type User struct {
	EmployeeId  int    `json:"employeeId"`
	Name        string `json:"name"`
	Mail        string `json:"mail"`
	Department  string `json:"department"`
	DateOfBirth string `json:"dateOfBirth"`
	Gender      string `json:"gender"`
	Type        string `json:"type"`
}

type Users []User
