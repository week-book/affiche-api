package repositorytest

import "github.com/week-book/affiche-api/internal/domain"

type EventRepository struct {
	CreateFunc func(event domain.Event) (string, error)
}

func (f *EventRepository) Create(event domain.Event) (string, error) {
	return f.CreateFunc(event)
}
