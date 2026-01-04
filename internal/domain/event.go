package domain

type Event struct {
	PhotoId string `json:"photo"`
	Text    string `json:"text"`
	Date    string `json:"date"`
}

type EventRepository interface {
	Create(event Event) (string, error)
}
