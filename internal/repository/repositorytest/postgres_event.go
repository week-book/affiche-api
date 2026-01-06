package repositorytest

import (
	"github.com/google/uuid"
	"github.com/week-book/affiche-api/internal/domain"
)

type EventRepository struct {
	CreateFunc  func(event domain.Event) (uuid.UUID, error)
	GetByIDFunc func(id uuid.UUID) (domain.Event, error)
}

func (r *EventRepository) Create(event domain.Event) (uuid.UUID, error) {
	return r.CreateFunc(event)
}

func (r *EventRepository) GetByID(id uuid.UUID) (domain.Event, error) {
	return r.GetByIDFunc(id)
}
