package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
)

var (
	dbUrl = os.Getenv("DATABASE_URL")
)

var connection *sql.DB

func Open() (*sql.DB, error) {
	var err error
	connection, err = sql.Open("postgres", dbUrl)
	if err != nil {
		return nil, err
	}

	return connection, nil
}

func Connection() *sql.DB {
	return connection
}
