package domain

type StoredFeedback struct {
	UserId       int
	FeelingId    int
	CategoryId   int
	RecipientId  int
	NewsId       int
	FeedbackNote string
}

type UserIdAndCategoryId struct {
	UserId          int
	PointCategoryId int
}

type UserIdAndPoints struct {
	EmployeeId string `json:"employeeId"`
	Points     int    `json:"points"`
}

type UserIdAndCategoryIds []UserIdAndCategoryId
