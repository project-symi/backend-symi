package usecase

type FeelingRepository interface {
	FindIdByName(name string) (id int, err error)
}
