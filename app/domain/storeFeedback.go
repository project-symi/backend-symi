package domain

type StoredFeedback struct {
	UserId           int
	FeelingId        int
	CategoryId       int
	RecipientId      int
	RecipientSlackId string
	NewsId           int
	FeedbackNote     string
}

type UserIdAndCategoryId struct {
	UserId          int
	PointCategoryId int
}

type StoredInfo struct {
	EmployeeId       string `json:"employeeId"`
	Points           int    `json:"points"`
	RecipientPoints  int    `json:"-"`
	RecipientSlackId string `json:"-"`
}

type UserIdAndCategoryIds []UserIdAndCategoryId
