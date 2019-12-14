package repository

type CategoryRepository interface {
	FindIdByName(name string) (id int, err error)
}
