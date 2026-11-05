package database

import (
	"database/sql"
	"embed"

	"github.com/pressly/goose/v3"
)

// ----- EXPLANATION -----
// Embedding our migrations inside of our binary:
// This way, we can use the migrations, when we run the app
// Without having to copy all of the migrations
// To our container when running

//go:embed migrations/*.sql
var migrationsFS embed.FS

func Migrate(db *sql.DB) error {
	goose.SetBaseFS(migrationsFS)

	if err := goose.SetDialect("mysql"); err != nil {
		return err
	}

	return goose.Up(db, "migrations")
}
