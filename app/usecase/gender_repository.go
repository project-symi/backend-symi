package usecase

type GenderRepository interface {
	GenderToId(gender string) (id int, err error)
}
