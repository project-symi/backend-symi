package database

import (
	"project-symi-backend/app/domain"
	"time"
)

type PointRepository struct {
	SqlHandler
}

func (repo *PointRepository) FindPointsByUserId(userId int) (points domain.Points, err error) {
	rows, err := repo.Query(`
		SELECT
			pc.name,
			pc.point,
			f.feedback_note,
			p.created_at
		FROM point_logs p
		JOIN point_categories pc ON pc.id = p.point_category_id
		JOIN feedbacks f ON f.id = p.feedback_id
		WHERE p.user_id = ?`, userId)
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var (
			categoryName string
			point        int
			feedbackNote string
			createdAt    string
		)
		if err := rows.Scan(
			&categoryName,
			&point,
			&feedbackNote,
			&createdAt); err != nil {
			continue
		}
		pointValue := domain.Point{
			CategoryName: categoryName,
			Point:        point,
			FeedbackNote: feedbackNote,
			CreatedAt:    createdAt,
		}
		points = append(points, pointValue)
	}
	return
}

func (repo *PointRepository) AddPointLog(userId int, pointCategorId int, feedbackId int, expireDate string) (insertedId int, err error) {
	result, err := repo.Execute(`
		INSERT INTO addPointLog (user_id, point_category_id, feedback_id, expire_date, created_at, modified_at)
		VALUES (?, ?, ?, ?, ?, ?)
	  `, userId, pointCategorId, feedbackId, expireDate, time.Now(), time.Now())
	insertedId64, err := result.LastInsertId()
	if err != nil {
		return
	}
	insertedId = int(insertedId64)
	return
}

func (repo *PointRepository) UpdateUserTotalPoint(userId int, point int) (success bool, err error) {
	result, err := repo.Execute(`
	UPDATE users u
	JOIN point_logs p ON p.user_id = u.id
	JOIN point_categories pc ON p.point_category_id = pc.id
	SET
	  u.total_points = u.total_points + pc.point,
	  u.modified_at = ?
	WHERE
	  p.id = ?
		`, time.Now(), userId)
	if err != nil {
		return
	}
	amountOfAffected64, err := result.RowsAffected()
	if err != nil {
		return
	}
	if amountOfAffected64 != 0 {
		success = true
		return
	}
	success = false
	return
}
