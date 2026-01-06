package repository

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/week-book/affiche-api/internal/domain"
)

var ErrEventNotFound = errors.New("event not found")

type PostgresEventRepository struct {
	db *sql.DB
}

func NewPostgresEventRepository(db *sql.DB) *PostgresEventRepository {
	return &PostgresEventRepository{db: db}
}

func (r *PostgresEventRepository) Create(event domain.Event) (uuid.UUID, error) {
	id := uuid.New()

	_, err := r.db.Exec(
		`INSERT INTO events (id, text, date, photo) VALUES ($1, $2, $3, $4)`,
		id,
		event.Text,
		event.Date,
		event.PhotoId,
	)
	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}

func (r *PostgresEventRepository) GetByID(id uuid.UUID) (domain.Event, error) {
	const query = `
		SELECT id, text, date, photo
		FROM events
		WHERE id = $1
	`

	row := r.db.QueryRow(query, id)

	var event domain.Event

	err := row.Scan(
		&event.ID,
		&event.Text,
		&event.Date,
		&event.PhotoId,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Event{}, ErrEventNotFound
		}
		return domain.Event{}, err
	}

	return event, nil
}
