package domain

type Reward struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Points int    `json:"points"`
	Url    string `json:"url"`
}

type Rewards []Reward
