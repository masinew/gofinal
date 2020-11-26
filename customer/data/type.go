package data

type Customer struct {
	Id int 			`json:"id"`
	Name string 	`json:"name"`
	Email string 	`json:"email"`
	Status string 	`json:"status"`
}

type Message struct {
	Message string	`json:"message"`
}
