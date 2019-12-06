package domain

type Feedback struct {
	// Id int
	UserId int
	// OpinionId   int
	// CategoryId  int
	FeedbackNote string
	// UpdatedAt string
	// CreatedAt string
}

type Feedbacks []Feedback
