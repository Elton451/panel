package main

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	config "github.com/elton451/panel-monorepo/services/api/internal/config"
	database "github.com/elton451/panel-monorepo/services/api/internal/database"
	"github.com/elton451/panel-monorepo/services/api/internal/models"
	_ "github.com/go-sql-driver/mysql"
)

type Application struct {
	Db     *sql.DB
	Orders *models.OrderModel
	logger *slog.Logger
}

func main() {
	config, err := config.Load()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	if err != nil {
		log.Fatal("Error while loading application config", err)
	}

	db, err := openDB(config.DbConnectString)
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	err = database.Migrate(db)
	if err != nil {
		fmt.Println("ERROR ON MIGRATION", err)
	}

	app := &Application{
		Db:     db,
		Orders: &models.OrderModel{Db: db},
		logger: logger,
	}

	fmt.Printf("Hello! Listenning on Port %s", config.Port)
	err = http.ListenAndServe(":"+config.Port, app.routes())
	if err != nil {
		log.Fatal("Error starting up server", err)
	}
}

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
