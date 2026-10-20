package main

import (
	"database/sql"
	"fmt"
	"net/http"

	utils "github.com/elton451/panel-monorepo/services/api/internals/utils"
	_ "github.com/go-sql-driver/mysql"
)

type Application struct {
	Db *sql.DB
}

func main() {
	envVars := utils.GetEnvVars()

	// Open DB Connection
	db, err := openDB(envVars.DbConnectString)
	if err != nil {
		fmt.Errorf("Error connecting to database", err)
	}

	app := &Application{
		Db: db,
	}

	http.ListenAndServe(":"+envVars.Port, app.routes())
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
