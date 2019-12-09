package domain

type Feedback struct {
	Id                  int    `json:"id"`
	EmployeeId          string `json:"employeeId"`
	Feeling             string `json:"feeling"`
	Seen                bool   `json:"status"`
	Category            string `json:"category"`
	RecipientEmployeeId string `json:"recipientId"`
	NewsId              int    `json:"newsId"`
	FeedbackNote        string `json:"note"`
	RecipientName       string `json:"name"`
	CreatedAT           string `json:"dateAdded"`
}

type Feedbacks []Feedback
