package domain

type User struct {
	EmployeeId  int
	Name        string
	Mail        string
	Department  string
	DateOfBirth string
	Gender      string
	Type        string
}

type Users []User
