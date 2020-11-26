package data

import (
	"github.com/masinew/gofinal/database"
)

const (
	customerTableQuery = `
		CREATE TABLE IF NOT EXISTS customer (
			id SERIAL PRIMARY KEY,
			name TEXT,
			email TEXT,
			status TEXT
		);
	`
)

func CreateTodosTable() error {
	conn := database.Connection()
	_, err := conn.Exec(customerTableQuery)
	if err != nil {
		return err
	}

	return nil
}

