package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

const (
	dbUrl = `postgres://tddwulka:UlftGmzu4h5oRstCi_pZhlcMXNB6EVye@rosie.db.elephantsql.com:5432/tddwulka`
	//dbUrl = os.Getenv("DB_URL")
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
