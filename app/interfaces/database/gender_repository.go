package database

type GenderRepository struct {
	SqlHandler
}

func (repo *GenderRepository) GenderToId(gender string) (id int, err error) {
	row, err := repo.Query("SELECT id from genders WHERE deleted = false AND gender = ?", gender)
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
