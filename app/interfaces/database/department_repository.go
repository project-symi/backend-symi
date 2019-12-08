package database

type DepartmentRepository struct {
	SqlHandler
}

func (repo *DepartmentRepository) DepartmentToId(department string) (id int, err error) {
	row, err := repo.Query("SELECT id from departments WHERE deleted = false AND name = ?", department)
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
