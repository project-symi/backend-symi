package usecase

type DepartmentRepository interface {
	DepartmentToId(department string) (id int, err error)
}
