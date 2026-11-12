package main

import (
	"database/sql"
	"log"

	"github.com/elton451/panel-monorepo/services/api/internal/config"
	databaseInternal "github.com/elton451/panel-monorepo/services/api/internal/database"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatal("Error while loading application config", err)
	}

	db, err := openDB(config.DbConnectString)
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	databaseInternal.Migrate(db)
}

// TODO: Refactor this to internal -> check other server main.go
func openDB(connector string) (*sql.DB, error) {
	db, err := sql.Open("mysql", connector)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
