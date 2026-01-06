package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/week-book/affiche-api/internal/domain"
	"github.com/week-book/affiche-api/internal/repository"
	"github.com/week-book/affiche-api/internal/service"
)

type EventHandler struct {
	service *service.EventService
}

func NewEventHandler(svc *service.EventService) *EventHandler {
	return &EventHandler{service: svc}
}

func (h *EventHandler) Create(w http.ResponseWriter, r *http.Request) {
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

func (h *EventHandler) GetEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		http.Error(w, service.ErrInvalidID.Error(), http.StatusBadRequest)
		return
	}

	event, err := h.service.GetByID(id)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidID):
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		case errors.Is(err, repository.ErrEventNotFound):
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
	}

	json, err := json.Marshal(event)
	if err != nil {
		http.Error(w, "Failed to convert response to json format", http.StatusInternalServerError)
	}

	w.Write(json)
}
