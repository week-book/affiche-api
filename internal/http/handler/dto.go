package handler

type EventResponse struct {
	ID      string `json:"id"`
	PhotoId string `json:"photo"`
	Text    string `json:"text"`
	Date    string `json:"date"`
}
