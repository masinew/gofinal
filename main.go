package main

import (
	"github.com/gin-gonic/gin"
	"github.com/masinew/gofinal/customer"
	"github.com/masinew/gofinal/customer/data"
	"github.com/masinew/gofinal/database"
	"log"
)

func main() {
	conn, connErr := database.Open()
	defer conn.Close()
	if connErr != nil {
		log.Fatalf("Database connection error: %s", connErr)
	}

	err := data.CreateTodosTable()
	if err != nil {
		log.Fatalf("Customer table creation error: %s", err)
	}

	r := gin.Default()
	r.GET("/customers", customer.FindAllCustomerHandler)
	r.GET("/customers/:id", customer.FindCustomerHandler)
	r.POST("/customers", customer.CreateCustomerHandler)
	r.PUT("/customers/:id", customer.UpdateCustomerHandler)
	r.DELETE("/customers/:id", customer.DeleteCustomerHandler)
	r.Run(":2009")
}
