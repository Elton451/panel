package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbConnectString string
	Port            string
}

func Load() (Config, error) {
	// Loading ENV
	err := godotenv.Load()
	if err != nil {
		return Config{}, err
	}

	dbConnectString := os.Getenv("DB_STRING")
	port := os.Getenv("PORT")

	config := Config{
		DbConnectString: dbConnectString,
		Port:            port,
	}

	return config, nil
}
