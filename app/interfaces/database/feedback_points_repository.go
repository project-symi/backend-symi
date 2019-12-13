package database

import (
	"project-symi-backend/app/domain"
	"time"
)

type FeedbackPointsRepository struct {
	SqlHandler
}

const (
	Poll = iota + 1
	SubmittedFeedback
	ReceivedPositiveFeedback
)

func (repo *FeedbackPointsRepository) StoreFeedbackAndUpdatePoints(storedFeedback domain.StoredFeedback, expireDate string) (points int, err error) {
	pointCategory := SubmittedFeedback
	tx, err := repo.Begin()
	if err != nil {
		return
	}
	feedbackId, err := addTxFeedback(tx, storedFeedback)
	if err != nil {
		tx.Rollback()
		return
	}
	err = addTxPointLog(tx, storedFeedback.UserId, pointCategory, feedbackId, expireDate)
	if err != nil {
		tx.Rollback()
		return
	}
	points, err = findTxPointsById(tx, pointCategory)
	err = updateTxUserTotalPoint(tx, storedFeedback.UserId, points, feedbackId)
	if err != nil {
		tx.Rollback()
		return
	}
	if storedFeedback.RecipientId != 0 {
		pointCategory = ReceivedPositiveFeedback
		err = addTxPointLog(tx, storedFeedback.RecipientId, pointCategory, feedbackId, expireDate)
		if err != nil {
			tx.Rollback()
			return
		}
		recipientPoints := 0
		recipientPoints, err = findTxPointsById(tx, pointCategory)
		if err != nil {
			tx.Rollback()
			return
		}
		err = updateTxUserTotalPoint(tx, storedFeedback.RecipientId, recipientPoints, feedbackId)
		if err != nil {
			tx.Rollback()
			return
		}
	}
	tx.Commit()
	return
}

func addTxFeedback(tx Tx, feedback domain.StoredFeedback) (insertedId int, err error) {
	result, err := tx.Execute(`
		INSERT INTO feedbacks (user_id, feeling_id, category_id, recipient_id, news_id, feedback_note, created_at, modified_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	  `, feedback.UserId, feedback.FeelingId, feedback.CategoryId, feedback.RecipientId, feedback.NewsId, feedback.FeedbackNote, time.Now(), time.Now())
	insertedId64, err := result.LastInsertId()
	if err != nil {
		return
	}
	insertedId = int(insertedId64)
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

func updateTxUserTotalPoint(tx Tx, userId int, point int, feedbackId int) (err error) {
	_, err = tx.Execute(`
	UPDATE users u
	JOIN point_logs p ON p.user_id = u.id
	JOIN point_categories pc ON p.point_category_id = pc.id
	SET
	  u.total_points = u.total_points + pc.point,
	  u.modified_at = ?
	WHERE
	  u.id = ?
	AND
	  p.feedback_id = ?
		`, time.Now(), userId, feedbackId)
	if err != nil {
		return
	}
	return
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
