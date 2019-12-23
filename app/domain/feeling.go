package domain

type FeelingQuery struct {
	Feeling string `form:"feeling" binding:"required,oneof=good meh sad"`
}
