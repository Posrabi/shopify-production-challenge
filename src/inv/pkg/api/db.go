package api

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// NewSQLDB returns a new sql database.
func NewSQLDB() (*sql.DB, error) {
	return sql.Open("postgres", getPostgresConnectionString())
}

func getPostgresConnectionString() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	// There is a lot more that should be done here like switching between dev/prod environment.
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
}
