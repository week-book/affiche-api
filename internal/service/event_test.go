package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/week-book/affiche-api/internal/domain"
	"github.com/week-book/affiche-api/internal/repository/repositorytest"
	"github.com/week-book/affiche-api/internal/service"
)

func setup() *service.EventService {
	repo := &repositorytest.EventRepository{
		CreateFunc: func(event domain.Event) (string, error) {
			return "test-id", nil
		},
	}

	svc := service.NewEventService(repo)

	return svc
}

func TestEventService_Create_ReturnsID(t *testing.T) {
	svc := setup()

	id, err := svc.Create(domain.Event{
		Text: "test event",
		Date: "2025-01-01",
	})

	assert.NoError(t, err)
	assert.NotEmpty(t, id)
}

func TestEventService_Create_EmptyText_ReturnsError(t *testing.T) {
	svc := setup()
	_, err := svc.Create(domain.Event{Text: ""})

	assert.ErrorIs(t, err, service.ErrEmptyText)
}
