package data

import "database/sql"

func Insert(db *sql.DB, customer *Customer) error {
	customer.Id = 101
	return nil
}

func Update(db *sql.DB, id int, customer *Customer) error {
	customer.Id = id
	return nil
}

func Delete(db *sql.DB, id int) error {
	return nil
}
