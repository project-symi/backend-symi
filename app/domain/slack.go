package domain

type Slack struct {
	Token string
	Url   string `json:"url"`
	Text  string `json:"text"`
}

type SlackBody struct {
	Channel string `json:"channel"`
	Text    string `json:"text"`
}
