package handler

import "github.com/google/uuid"

type EventResponse struct {
	ID      uuid.UUID `json:"id"`
	PhotoId string    `json:"photo"`
	Text    string    `json:"text"`
	Date    string    `json:"date"`
}
