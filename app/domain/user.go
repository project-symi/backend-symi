package domain

type User struct {
	EmployeeId   string `json:"employeeId"`
	Name         string `json:"name"`
	Password     string
	Mail         string `json:"email"`
	Department   string `json:"department"`
	DateOfBirth  string `json:"dateOfBirth"`
	TotalPoints  int    `json:"totalPoints"`
	Gender       string `json:"gender"`
	Permission   string `json:"permission"`
	CurrentToken string
}

type Users []User
