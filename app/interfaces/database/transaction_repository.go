package database

import (
	"project-symi-backend/app/domain"
)

type TransactionRepository struct {
	SqlHandler
}

const (
	SubmittedFeedback = iota + 1
	ReceivedPositiveFeedback
)

func (repo *TransactionRepository) StoreFeedbackAndUpdatePoints(storedFeedback domain.StoredFeedback, expireDate string) (points int, recipientPoints int, err error) {
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
		recipientPoints = 0
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
