package domain

type Feedback struct {
	EmployeeId          string `json:"employeeId"`
	Feeling             string `json:"feeling"`
	Seen                bool   `json:"status"`
	Category            string `json:"category"`
	RecipientEmployeeId string `json:"recipientId"`
	NewsId              int    `json:"newsId"`
	FeedbackNote        string `json:"note"`
}

type Feedbacks []Feedback
