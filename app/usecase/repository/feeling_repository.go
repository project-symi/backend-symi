package repository

type FeelingRepository interface {
	FindIdByName(name string) (id int, err error)
}
