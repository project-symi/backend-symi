package domain

type Point struct {
	CategoryName          string `json:"categoryName"`
	Point                 int    `json:"points"`
	FeedbackCategory      string `json:"feedbackCategory"`
	FeedbackNote          string `json:"feedbackNote"`
	FeedbackRecipientName string `json:"feedbackRecipientName"`
	FeedbackNewsTitle     string `json:"feedbackNewsTitle"`
	CreatedAt             string `json:"date"`
}

type Points []Point
