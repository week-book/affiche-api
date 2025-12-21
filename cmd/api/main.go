package main

import (
	"log"
	"net/http"
	"os"

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

	repo := repository.NewPostgresEventRepository(dbConn)
	svc := service.NewEventService(repo)
	h := handler.NewEventHandler(svc)

	mux := http.NewServeMux()
	mux.HandleFunc("/events", h.Create)

	log.Println("server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
