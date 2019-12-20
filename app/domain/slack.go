package domain

type Slack struct {
	Token string
	Url   string
	Text  string
}

type SlackBody struct {
	Channel string `json:"channel"`
	Text    string `json:"text"`
}
