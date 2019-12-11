package domain

type Point struct {
	CategoryName string `json:"categoryName"`
	Point        int    `json:"points"`
	FeedbackNote string `json:"feedbackNote"`
	CreatedAt    string `json:"date"`
}

type Points []Point
