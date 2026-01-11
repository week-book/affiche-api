package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/week-book/affiche-api/internal/db"
	"github.com/week-book/affiche-api/internal/http/handler"
	"github.com/week-book/affiche-api/internal/repository"
	"github.com/week-book/affiche-api/internal/service"
)

func main() {
	dbConn, err := db.NewPostgres(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()
	log.Println("database connected")

	health := handler.NewHealthHandler(dbConn)
	go func() {
		t := time.NewTicker(5 * time.Second)
		defer t.Stop()
		for range t.C {
			health.Touch()
		}
	}()

	repo := repository.NewPostgresEventRepository(dbConn)
	svc := service.NewEventService(repo)
	h := handler.NewEventHandler(svc)
	r := mux.NewRouter()

	r.HandleFunc("/events", h.Create).Methods(http.MethodPost)
	r.HandleFunc("/events/{id}", h.GetEvent).Methods(http.MethodGet)
	r.HandleFunc("/healthz", health.Liveness).Methods(http.MethodGet)
	r.HandleFunc("/readyz", health.Readiness).Methods(http.MethodGet)

	log.Println("server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
