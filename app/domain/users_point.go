package domain

type UserWithPoint struct {
	Id         int
	Name       string
	Point      int
	Department string
	Gender     string
}

type UsersWithPoint []UserWithPoint
