package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// All the Env
type Env struct {
	DbConnectString string
	Port            string
}

func GetEnvVars() Env {
	// Loading ENV
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbConnectString := os.Getenv("DB_STRING")
	port := os.Getenv("PORT")

	return Env{
		DbConnectString: dbConnectString,
		Port:            port,
	}
}
