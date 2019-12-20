package domain

type Feedback struct {
	FeedbackId          int    `json:"feedbackId"`
	EmployeeId          string `json:"employeeId"`
	Feeling             string `json:"feeling"`
	Seen                bool   `json:"status"`
	Category            string `json:"category"`
	RecipientEmployeeId string `json:"recipientId"`
	NewsId              int    `json:"newsId"`
	FeedbackNote        string `json:"note"`
	CreatedAt           string `json:"dateAdded"`
}

type FeedbackForCEO struct {
	FeedbackId          int    `json:"feedbackId"`
	EmployeeId          string `json:"employeeId"`
	Feeling             string `json:"feeling"`
	Department          string `json:"department"`
	Seen                bool   `json:"status"`
	Category            string `json:"category"`
	RecipientEmployeeId string `json:"recipientId"`
	NewsId              int    `json:"newsId"`
	FeedbackNote        string `json:"note"`
	CreatedAt           string `json:"dateAdded"`
}

type Feedbacks []Feedback
type FeedbacksForCEO []FeedbackForCEO

type FeedbackStore struct {
	EmployeeId          string `json:"employeeId"`
	Feeling             string `json:"feeling"`
	Category            string `json:"category"`
	RecipientEmployeeId string `json:"recipientId"`
	NewsId              int    `json:"newsId"`
	FeedbackNote        string `json:"note"`
}

type FeedbackEmployee struct {
	Id            int    `json:"id"`
	Feeling       string `json:"feeling"`
	Seen          bool   `json:"status"`
	Category      string `json:"category"`
	NewsId        int    `json:"newsId"`
	FeedbackNote  string `json:"note"`
	RecipientName string `json:"name"`
	CreatedAT     string `json:"dateAdded"`
}

type FeedbackEmployees []FeedbackEmployee
