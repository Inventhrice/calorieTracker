package middlewares

import (
	"database/sql"
	"fmt"
	"io/fs"
	"os"
	"slices"
	"strings"

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
	if err := runMigrations(); err != nil {
		return err
	}

	return nil

}

func runMigrations() error {
	var migrationVersion string
	migrationsDir := "/app/migrations/"
	if files, err := os.ReadDir(migrationsDir); err != nil {
		return err
	} else {
		slices.SortFunc(files, func(a, b fs.DirEntry) int {
			return strings.Compare(a.Name(), b.Name())
		})
		if err := Database.Get(&migrationVersion, "SELECT value FROM migrations WHERE key=\"migrations\""); err != nil {
			if err == sql.ErrNoRows {
				return executeMigrationScript(files[0].Name())
			} else {
				return err
			}

		} else {
			index := slices.IndexFunc(files, func(a fs.DirEntry) bool {
				return strings.Contains(a.Name(), migrationVersion)
			})
			for i := index + 1; i < len(files); i++ {
				return executeMigrationScript(files[i].Name())
			}
			return nil
		}
	}
}

func executeMigrationScript(filename string) error {
	if data, err := os.ReadFile(filename); err != nil {
		return err
	} else {
		tx, _ := Database.Begin()
		if _, err := tx.Exec(string(data)); err != nil {
			tx.Rollback()
			return err
		}
		tx.Commit()
		return nil
	}
}
