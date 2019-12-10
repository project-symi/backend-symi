package database

type CategoryRepository struct {
	SqlHandler
}

func (repo *CategoryRepository) FindIdByName(name string) (id int, err error) {
	row, err := repo.Query(`SELECT id FROM categories WHERE name = ?`, name)
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
