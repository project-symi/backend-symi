package usecase

type PermissionRepository interface {
	PermissionToId(permission string) (id int, err error)
}
