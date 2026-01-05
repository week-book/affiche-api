package repository_test

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"

	"github.com/week-book/affiche-api/internal/config"
	"github.com/week-book/affiche-api/internal/db"
	"github.com/week-book/affiche-api/internal/domain"
	"github.com/week-book/affiche-api/internal/repository"
)

func TestMain(m *testing.M) {
	config.Load()
	code := m.Run()
	os.Exit(code)
}

func newTestDB(t *testing.T) *sql.DB {
	t.Helper()

	dbConn, err := db.NewPostgres(
		"localhost",
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		dbConn.Close()
	})

	return dbConn
}

func TestPostgresEventRepository_Create(t *testing.T) {
	db := newTestDB(t)

	db.Exec("DELETE FROM events")
	repo := repository.NewPostgresEventRepository(db)

	event := domain.Event{
		PhotoId: "1",
		Text:    "test event",
		Date:    "2025-01-01",
	}

	id, err := repo.Create(event)

	assert.NoError(t, err)
	assert.NotEmpty(t, id)

	var text string
	err = db.QueryRow(
		`SELECT text FROM events WHERE id = $1`,
		id,
	).Scan(&text)

	assert.NoError(t, err)
	assert.Equal(t, "test event", text)
}
