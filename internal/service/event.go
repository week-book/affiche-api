package service

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/week-book/affiche-api/internal/domain"
)

type EventService struct {
	repo domain.EventRepository
}

var (
	ErrEmptyText  = errors.New("text is empty")
	ErrEmptyPhoto = errors.New("photo is empty")
	ErrInvalidID  = errors.New("id is invalid")
)

func NewEventService(repo domain.EventRepository) *EventService {
	return &EventService{repo: repo}
}

func (s *EventService) Create(eventInput domain.Event) (uuid.UUID, error) {

	if strings.TrimSpace(eventInput.Text) == "" {
		return uuid.UUID{}, ErrEmptyText
	}

	if strings.TrimSpace(eventInput.PhotoId) == "" {
		return uuid.UUID{}, ErrEmptyPhoto
	}
	return s.repo.Create(eventInput)
}

func (s *EventService) GetByID(id string) (domain.Event, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return domain.Event{}, ErrInvalidID
	}

	return s.repo.GetByID(parsedID)
}
