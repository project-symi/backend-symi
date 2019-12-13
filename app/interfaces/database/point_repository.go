package database

import (
	"project-symi-backend/app/domain"
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
