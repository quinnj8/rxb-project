package models

type Film struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Rating   string `json:"rating"`
	Category string `json:"category"`
}

type Films struct {
	Films []Film `json:"films"`
}
