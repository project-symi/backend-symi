package domain

type NewsItem struct {
	NewsItemId  int    `json:"newsId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	PhotoLink   string `json:"photo"`
	Hidden      bool   `json:"status"`
	CreatedAt   string `json:"postedOn"`
	ModifiedAt  string `json:"modiefiedOn"`
}

type News []NewsItem
