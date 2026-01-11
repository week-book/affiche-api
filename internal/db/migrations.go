package db

import (
	"context"
	"database/sql"
	"errors"
)

var ErrMigrationsDirty = errors.New("database migrations are dirty")
var ErrMigrationsNotApplied = errors.New("required migrations are not applied")

func CheckMigrations(
	ctx context.Context,
	db *sql.DB,
	expectedVersion int64,
) error {
	var version int64
	var dirty bool

	err := db.QueryRowContext(
		ctx,
		`SELECT version, dirty FROM schema_migrations LIMIT 1`,
	).Scan(&version, &dirty)

	if err != nil {
		return err
	}

	if dirty {
		return ErrMigrationsDirty
	}

	if version < expectedVersion {
		return ErrMigrationsNotApplied
	}

	return nil
}
