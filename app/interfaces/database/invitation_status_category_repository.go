package database

type InvitationStatusCategoryRepository struct {
	SqlHandler
}

func (repo *InvitationStatusCategoryRepository) FindKeyIdByStatus(status string) (id int, err error) {
	rows, err := repo.Query("SELECT id FROM invitation_status_categories WHERE status = ?", status)
	defer rows.Close()
	if err != nil {
		return
	}
	rows.Next()
	if err = rows.Scan(&id); err != nil {
		return
	}
	return
}
