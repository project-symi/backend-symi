package domain

type Reward struct {
	Id     int    `json:"id" binding:"required,numeric"`
	Name   string `json:"name" binding:"required,max=50,excludesall=!()#@{}?"`
	Points int    `json:"points" binding:"required,numeric"`
	Url    string `json:"url" binding:"required,url"`
}

type Rewards []Reward
