package domain

type Event struct {
	Text string `json:"text"`
	Date string `json:"date"`
}

type EventRepository interface {
	Create(event Event) (string, error)
}
