package handler_test

import (
	"bytes"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/week-book/affiche-api/internal/domain"
	"github.com/week-book/affiche-api/internal/http/handler"
	"github.com/week-book/affiche-api/internal/repository/repositorytest"
	"github.com/week-book/affiche-api/internal/service"
)

func setup() *handler.EventHandler {
	repo := &repositorytest.EventRepository{
		CreateFunc: func(event domain.Event) (string, error) {
			return "test-id", nil
		},
	}

	svc := service.NewEventService(repo)
	eventHandler := handler.NewEventHandler(svc)

	return eventHandler
}

func TestCreateEvent_Returns201AndID(t *testing.T) {
	json := `{"photo": "1", "text": "test event", "date": "2025-01-01"}`
	r := httptest.NewRequest("POST", "/events", bytes.NewBuffer([]byte(json)))
	w := httptest.NewRecorder()

	eventHandler := setup()
	eventHandler.Create(w, r)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	respCode := resp.StatusCode

	assert.Equal(t, 201, respCode)
	assert.Contains(t, string(body), "id")
}

func TestCreateEvent_MethodNotAllowed(t *testing.T) {
	r := httptest.NewRequest("GET", "/events", nil)
	w := httptest.NewRecorder()

	eventHandler := setup()
	eventHandler.Create(w, r)

	assert.Equal(t, 405, w.Result().StatusCode)
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
