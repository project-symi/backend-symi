package database

type FeelingRepository struct {
	SqlHandler
}

func (repo *FeelingRepository) FindIdByName(name string) (id int, err error) {
	row, err := repo.Query(`SELECT id FROM feelings WHERE name = ?`, name)
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
