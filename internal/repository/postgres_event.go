package repository

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/week-book/affiche-api/internal/domain"
)

type PostgresEventRepository struct {
	db *sql.DB
}

func NewPostgresEventRepository(db *sql.DB) *PostgresEventRepository {
	return &PostgresEventRepository{db: db}
}

func (r *PostgresEventRepository) Create(event domain.Event) (string, error) {
	id := uuid.New().String()

	_, err := r.db.Exec(
		`INSERT INTO events (id, text, date, photo) VALUES ($1, $2, $3, $4)`,
		id,
		event.Text,
		event.Date,
		event.PhotoId,
	)
	if err != nil {
		return "", err
	}

	return id, nil
}
