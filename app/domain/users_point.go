package domain

type UserWithPoints struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Points      int    `json:"points"`
	DateOfBirth string `json:"dateOfBirth"`
	Department  string `json:"department"`
	Gender      string `json:"gender"`
}

type UsersWithPoints []UserWithPoints
