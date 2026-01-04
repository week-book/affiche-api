package handler

import (
	"encoding/json"
	"net/http"

	"github.com/week-book/affiche-api/internal/domain"
	"github.com/week-book/affiche-api/internal/service"
)

type EventHandler struct {
	service *service.EventService
}

func NewEventHandler(svc *service.EventService) *EventHandler {
	return &EventHandler{service: svc}
}

func (h *EventHandler) Create(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "An incorrect request method was selected", http.StatusMethodNotAllowed)
		return
	}

	event := domain.Event{}
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	if err := decoder.Decode(&event); err != nil {
		http.Error(w, "Invalid json", http.StatusBadRequest)
		return
	}

	id, err := h.service.Create(event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	eventResponse := EventResponse{
		ID:      id,
		PhotoId: event.PhotoId,
		Text:    event.Text,
		Date:    event.Date,
	}

	jsonResponse, err := json.Marshal(eventResponse)
	if err != nil {
		http.Error(w, "Failed to convert response to json format", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}
