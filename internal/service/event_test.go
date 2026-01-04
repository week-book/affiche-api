package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/week-book/affiche-api/internal/domain"
	"github.com/week-book/affiche-api/internal/repository/repositorytest"
	"github.com/week-book/affiche-api/internal/service"
)

func setup() (*service.EventService, domain.Event) {
	repo := &repositorytest.EventRepository{
		CreateFunc: func(event domain.Event) (string, error) {
			return "test-id", nil
		},
	}

	svc := service.NewEventService(repo)

	domainEvent := domain.Event{
		PhotoId: "1",
		Text:    "test event",
		Date:    "2025-01-01",
	}
	return svc, domainEvent
}

func TestEventService_Create_ReturnsID(t *testing.T) {
	svc, de := setup()

	id, err := svc.Create(de)

	assert.NoError(t, err)
	assert.NotEmpty(t, id)
}

func TestEventService_Create_EmptyText_ReturnsError(t *testing.T) {
	svc, de := setup()

	de.Text = ""
	_, err := svc.Create(de)

	assert.ErrorIs(t, err, service.ErrEmptyText)
}

func TestEventService_Create_EmptyPhoto_ReturnsError(t *testing.T) {
	svc, de := setup()
	de.PhotoId = ""
	_, err := svc.Create(de)

	assert.ErrorIs(t, err, service.ErrEmptyPhoto)
}
