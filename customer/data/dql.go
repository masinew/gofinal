package data

import "database/sql"

func FindOne(db *sql.DB, id int) Customer {
	customer := Customer{
		Id: id,
		Name: "Champ",
		Email: "Email",
		Status: "Active",
	}

	return customer
}

func FindAll(db *sql.DB) []Customer {
	customers := []Customer{
		{
			Id: 1,
			Name: "Champ1",
			Email: "Email1",
			Status: "Active1",
		},
		{
			Id: 2,
			Name: "Champ2",
			Email: "Email2",
			Status: "Active2",
		},
	}

	return customers
}
