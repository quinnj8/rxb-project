package Models

type Film struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type Films struct {
	Films []Film `json:"films"`
}
