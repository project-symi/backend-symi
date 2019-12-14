package database

type PointCategoryRepository struct {
	SqlHandler
}

func findTxPointsById(tx Tx, id int) (point int, err error) {
	row, err := tx.Query(`
	SELECT
		point
	FROM point_categories
	WHERE deleted = false
	AND id = ?
	`, id)
	defer row.Close()
	if err != nil {
		return
	}
	row.Next()
	if err = row.Scan(
		&point); err != nil {
		return
	}
	return
}
