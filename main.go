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
	r.GET("/customer", customer.FindAllCustomerHandler)
	r.GET("/customer/:id", customer.FindCustomerHandler)
	r.POST("/customer", customer.CreateCustomerHandler)
	r.PUT("/customer/:id", customer.UpdateCustomerHandler)
	r.DELETE("/customer/:id")
	r.Run(":2009")
}
