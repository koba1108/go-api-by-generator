package config

import (
	"fmt"
	"os"
)

func DBConnectURL() string {
	dbDriver := os.Getenv("DB_DRIVER")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	return fmt.Sprintf("%s://%s:%s@%s/%s?parseTime=true", dbDriver, dbUser, dbPass, dbHost, dbName)
}
