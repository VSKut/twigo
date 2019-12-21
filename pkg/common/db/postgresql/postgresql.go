package postgresql

import (
	"database/sql"

	// Postgres driver
	_ "github.com/lib/pq"
)

// ConnectDB receives postgres connection string and returns *sql.DB with connection
func ConnectDB(connString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
