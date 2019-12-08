package database

type PermissionRepository struct {
	SqlHandler
}

func (repo *PermissionRepository) PermissionToId(permission string) (id int, err error) {
	row, err := repo.Query("SELECT id from permissions WHERE deleted = false AND name = ?", permission)
	defer row.Close()
	if err != nil {
		return
	}
	row.Next()
	if err = row.Scan(&id); err != nil {
		return
	}
	return
}
