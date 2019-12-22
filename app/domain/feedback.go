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
	EmployeeId          string `json:"employeeId" binding:"required,max=20,alphanum"`
	Feeling             string `json:"feeling" binding:"required,oneof=good meh sad"`
	Category            string `json:"category" binding:"required,max=20,excludesall=!()#@{}?"`
	RecipientEmployeeId string `json:"recipientId" binding:"max=20,alphanum"`
	NewsId              int    `json:"newsId" binding:"numeric"`
	FeedbackNote        string `json:"note" binding:"required"`
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
