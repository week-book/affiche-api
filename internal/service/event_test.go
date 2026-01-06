package service_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/week-book/affiche-api/internal/domain"
	"github.com/week-book/affiche-api/internal/repository/repositorytest"
	"github.com/week-book/affiche-api/internal/service"
)

func setup() *service.EventService {
	repo := &repositorytest.EventRepository{
		CreateFunc: func(event domain.Event) (uuid.UUID, error) {
			return uuid.MustParse("e16136c6-ae71-40df-983c-62119c5edb70"), nil
		},
		GetByIDFunc: func(id uuid.UUID) (domain.Event, error) {
			return domain.Event{
				ID:      id,
				PhotoId: "1",
				Text:    "test event",
				Date:    "2025-01-01",
			}, nil
		},
	}

	return service.NewEventService(repo)
}

func TestEventService_Create_ReturnsID(t *testing.T) {
	svc := setup()

	de := domain.Event{
		PhotoId: "1",
		Text:    "test event",
		Date:    "2025-01-01",
	}
	id, err := svc.Create(de)

	assert.NoError(t, err)
	assert.NotEmpty(t, id)
}

func TestEventService_Create_EmptyText_ReturnsError(t *testing.T) {
	svc := setup()

	de := domain.Event{
		PhotoId: "1",
		Text:    "",
		Date:    "2025-01-01",
	}
	de.Text = ""
	_, err := svc.Create(de)

	assert.ErrorIs(t, err, service.ErrEmptyText)
}

func TestEventService_Create_EmptyPhoto_ReturnsError(t *testing.T) {
	svc := setup()
	de := domain.Event{
		PhotoId: "",
		Text:    "test event",
		Date:    "2025-01-01",
	}
	de.PhotoId = ""
	_, err := svc.Create(de)

	assert.ErrorIs(t, err, service.ErrEmptyPhoto)
}
