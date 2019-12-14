package database

import (
	"project-symi-backend/app/domain"
	"time"
)

type PointsRepository struct {
	SqlHandler
}

func (repo *PointsRepository) FindPointsByUserId(userId int) (points domain.Points, err error) {
	rows, err := repo.Query(`
		SELECT
			pc.name,
			pc.point,
			f.feedback_note,
			p.created_at
		FROM point_logs p
		JOIN feedbacks f ON f.id = p.feedback_id
		JOIN point_categories pc ON pc.id = p.point_category_id
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

func addTxPointLog(tx Tx, userId int, pointCategorId int, feedbackId int, expireDate string) (err error) {
	_, err = tx.Execute(`
		INSERT INTO point_logs (user_id, point_category_id, feedback_id, expire_date, created_at, modified_at)
		VALUES (?, ?, ?, ?, ?, ?)
	  `, userId, pointCategorId, feedbackId, expireDate, time.Now(), time.Now())
	if err != nil {
		return
	}
	return
}
