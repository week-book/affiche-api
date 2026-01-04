package service

import (
	"errors"
	"strings"

	"github.com/week-book/affiche-api/internal/domain"
)

type EventService struct {
	repo domain.EventRepository
}

var ErrEmptyText = errors.New("text is empty")
var ErrEmptyPhoto = errors.New("photo is empty")

func NewEventService(repo domain.EventRepository) *EventService {
	return &EventService{repo: repo}
}

func (s *EventService) Create(eventInput domain.Event) (string, error) {

	if strings.TrimSpace(eventInput.Text) == "" {
		return "", ErrEmptyText
	}

	if strings.TrimSpace(eventInput.PhotoId) == "" {
		return "", ErrEmptyPhoto
	}
	return s.repo.Create(eventInput)
}
