package middlewares

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
)

var (
	Database *sqlx.DB
)

func InitDB() error {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DBNAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	// Format the connection URL and output to a string
	connectionURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
	fmt.Println(connectionURL)
	database, ConnectionError := sqlx.Open("mysql", connectionURL)

	if ConnectionError != nil {
		return ConnectionError
	}
	if err := database.Ping(); err != nil {
		return err
	}
	Database = database

	fmt.Println("Connected to database.")

	return nil

}
