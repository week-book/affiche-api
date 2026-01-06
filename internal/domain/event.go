package domain

import "github.com/google/uuid"

type Event struct {
	ID      uuid.UUID `json:"id"`
	PhotoId string    `json:"photo"`
	Text    string    `json:"text"`
	Date    string    `json:"date"`
}

type EventRepository interface {
	Create(event Event) (uuid.UUID, error)
	GetByID(parsedID uuid.UUID) (Event, error)
}
