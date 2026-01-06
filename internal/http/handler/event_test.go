package handler_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/week-book/affiche-api/internal/domain"
	"github.com/week-book/affiche-api/internal/http/handler"
	"github.com/week-book/affiche-api/internal/repository"
	"github.com/week-book/affiche-api/internal/repository/repositorytest"
	"github.com/week-book/affiche-api/internal/service"
)

func setup() *handler.EventHandler {
	repo := &repositorytest.EventRepository{
		CreateFunc: func(event domain.Event) (uuid.UUID, error) {
			return uuid.MustParse("e16136c6-ae71-40df-983c-62119c5edb70"), nil
		},
		GetByIDFunc: func(id uuid.UUID) (domain.Event, error) {
			existing := uuid.MustParse("e16136c6-ae71-40df-983c-62119c5edb70")
			if id == existing {
				return domain.Event{
					ID:      id,
					PhotoId: "1",
					Text:    "test event",
					Date:    "2025-01-01",
				}, nil
			}
			return domain.Event{}, repository.ErrEventNotFound
		},
	}
	svc := service.NewEventService(repo)
	eventHandler := handler.NewEventHandler(svc)

	return eventHandler
}

func createEvent(t *testing.T) handler.EventResponse {
	testJson := `{"photo": "1", "text": "test event", "date": "2025-01-01"}`
	r := httptest.NewRequest("POST", "/events", bytes.NewBuffer([]byte(testJson)))
	w := httptest.NewRecorder()

	h := setup()
	h.Create(w, r)

	require.Equal(t, http.StatusCreated, w.Result().StatusCode)

	var event handler.EventResponse
	require.NoError(t, json.NewDecoder(w.Body).Decode(&event))
	return event
}

func TestCreateEvent_Returns201AndResData(t *testing.T) {
	event := createEvent(t)
	assert.Equal(t, "1", event.PhotoId)
	assert.Equal(t, "test event", event.Text)
	assert.Equal(t, "2025-01-01", event.Date)

	assert.NotEmpty(t, event.ID)
}

func TestEventHandler_Create_EmptyText_Returns400(t *testing.T) {
	json := `{"text": "", "date": "2025-01-01"}`
	r := httptest.NewRequest("POST", "/events", bytes.NewBuffer([]byte(json)))
	w := httptest.NewRecorder()

	eventHandler := setup()
	eventHandler.Create(w, r)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, 400, w.Result().StatusCode)
	assert.Equal(t, string(body), service.ErrEmptyText.Error()+"\n")
}

func TestEventHandler_Create_EmptyPhoto_Returns400(t *testing.T) {
	json := `{"photo":"", "text": "some text", "date": "2025-01-01"}`
	r := httptest.NewRequest("POST", "/events", bytes.NewBuffer([]byte(json)))
	w := httptest.NewRecorder()

	eventHandler := setup()
	eventHandler.Create(w, r)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, 400, w.Result().StatusCode)
	assert.Equal(t, string(body), service.ErrEmptyPhoto.Error()+"\n")
}

func TestEventHandler_Create_InvalidJSON_Returns400(t *testing.T) {
	json := `{"text": "test event", "date": "2025-01-01"`
	r := httptest.NewRequest("POST", "/events", bytes.NewBuffer([]byte(json)))
	w := httptest.NewRecorder()

	eventHandler := setup()
	eventHandler.Create(w, r)

	assert.Equal(t, 400, w.Result().StatusCode)
}

func TestGetEvent_Returns200AndResData(t *testing.T) {
	event := createEvent(t)
	handle := "/events" + event.ID.String()
	r := httptest.NewRequest("GET", handle, nil)
	r = mux.SetURLVars(r, map[string]string{
		"id": event.ID.String(),
	})
	w := httptest.NewRecorder()

	eventHandler := setup()
	eventHandler.GetEvent(w, r)

	resp := w.Result()
	respCode := resp.StatusCode

	domainEvent := domain.Event{}
	decoder := json.NewDecoder(resp.Body)
	defer resp.Body.Close()
	if err := decoder.Decode(&domainEvent); err != nil {
		t.Errorf("Invalid json")
		return
	}

	assert.Equal(t, 200, respCode)

	assert.Equal(t, "1", event.PhotoId)
	assert.Equal(t, "test event", event.Text)
	assert.Equal(t, "2025-01-01", event.Date)

	assert.NotEmpty(t, event.ID)
}

func TestGetEvent_IdIsInvalid_Returns400(t *testing.T) {
	r := httptest.NewRequest("GET", "/events/abc", nil)
	w := httptest.NewRecorder()

	eventHandler := setup()
	eventHandler.GetEvent(w, r)

	resp := w.Result()
	respCode := resp.StatusCode
	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, 400, respCode)
	assert.Equal(t, service.ErrInvalidID.Error()+"\n", string(body))
}

func TestGetEvent_IdNotExist_Returns404(t *testing.T) {
	r := httptest.NewRequest("GET", "/events/e16136c6-ae71-40df-983c-62119c5edb71", nil)
	r = mux.SetURLVars(r, map[string]string{
		"id": "e16136c6-ae71-40df-983c-62119c5edb71",
	})
	w := httptest.NewRecorder()

	eventHandler := setup()
	eventHandler.GetEvent(w, r)

	resp := w.Result()
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, 404, resp.StatusCode)
	assert.Equal(t, repository.ErrEventNotFound.Error()+"\n", string(body))
}
