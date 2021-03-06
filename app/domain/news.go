package domain

type NewsItem struct {
	NewsItemId  int    `json:"newsId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	PhotoLink   string `json:"photo"`
	Hidden      bool   `json:"status"`
	CreatedAt   string `json:"postedOn"`
	ModifiedAt  string `json:"modifiedOn"`
}

type NewsPost struct {
	Title       string `json:"title" binding:"required,max=150"`
	Description string `json:"description" binding:"required"`
	PhotoLink   string `json:"photo" binding:"required,url"`
}

type News []NewsItem
