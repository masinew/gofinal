package data

import "database/sql"

func FindOne(db *sql.DB, id int) (Customer, error) {
	var customer Customer
	row := db.QueryRow("SELECT id, name, email, status FROM customer WHERE id = $1", id)
	err := row.Scan(&customer.Id, &customer.Name, &customer.Email, &customer.Status)
	if err != nil {
		return customer, err
	}

	return customer, nil
}

func FindAll(db *sql.DB) ([]Customer, error) {
	var customers []Customer
	result, err := db.Query("SELECT id, name, email, status FROM customer;")
	if err != nil {
		return customers, err
	}

	for result.Next() {
		var customer Customer
		sErr := result.Scan(&customer.Id, &customer.Name, &customer.Email, &customer.Status)
		if sErr != nil {
			return customers, sErr
		}

		customers = append(customers, customer)
	}

	return customers, nil
}
