package data

import (
	"database/sql"
	"errors"
)

func Insert(db *sql.DB, customer *Customer) error {
	row := db.QueryRow("INSERT INTO customer(name, email, status) VALUES($1, $2, $3) RETURNING id;",
		customer.Name, customer.Email, customer.Status)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return errors.New("insert error")
	}

	customer.Id = id
	return nil
}

func Update(db *sql.DB, id int, customer *Customer) error {
	_, err := db.Exec("UPDATE customer SET name = $2, email = $3, status = $4 WHERE id = $1;",
		id, customer.Name, customer.Email, customer.Status)
	if err != nil {
		return errors.New("update error")
	}

	customer.Id = id
	return nil
}

func Delete(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM customer WHERE id = $1", id)
	if err != nil {
		return errors.New("delete error")
	}

	return nil
}
