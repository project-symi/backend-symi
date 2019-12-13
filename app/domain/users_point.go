package domain

type UserWithPoint struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Point       int    `json:"point"`
	DateOfBirth string `json:"dateOfBirth"`
	Department  string `json:"department"`
	Gender      string `json:"gender"`
}

type UsersWithPoint []UserWithPoint
