package middlewares

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"slices"
	"strings"

	"github.com/jmoiron/sqlx"
)

const migrationsDir = "/app/migrations/"

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
	fmt.Println("Running migrations...")
	var migrationVersion string
	if files, err := os.ReadDir(migrationsDir); err != nil {
		return err
	} else {
		slices.SortFunc(files, func(a, b fs.DirEntry) int {
			return strings.Compare(a.Name(), b.Name())
		})
		if err := Database.Get(&migrationVersion, "SELECT value FROM metadata WHERE `key`=?", "migrations"); err != nil {
			fmt.Println("No value found in metadata table, running the inital migration.")
			if err := executeMigrationScript(files[0].Name()); err != nil {
				return err
			}
			migrationVersion = "1"
		}
		index := slices.IndexFunc(files, func(a fs.DirEntry) bool {
			return strings.Contains(a.Name(), migrationVersion)
		})
		for i := index + 1; i < len(files); i++ {
			fmt.Println("Running migration " + files[i].Name())
			if err := executeMigrationScript(files[i].Name()); err != nil {
				return err
			}
		}
		fmt.Println("Migrations complete with no errors!")
		return nil
	}
}

func executeMigrationScript(filename string) error {
	if data, err := os.ReadFile(migrationsDir + filename); err != nil {
		return err
	} else {
		tx, _ := Database.Begin()
		for _, cmd := range strings.Split(string(data), ";") {
			if cmd != "" {
				if _, err := tx.Exec(cmd); err != nil {
					tx.Rollback()
					return errors.New("Query " + cmd + " failed with error: " + err.Error())
				}
			}

		}
		tx.Commit()
		return nil
	}
}
