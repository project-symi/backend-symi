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
