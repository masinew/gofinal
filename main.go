package main

import (
	"github.com/gin-gonic/gin"
	"github.com/masinew/gofinal/customer"
	db "github.com/masinew/gofinal/database"
	"log"
)

func main() {
	conn, err := db.Open()
	defer conn.Close()
	if err != nil {
		log.Fatalf("Database Connection error: %s", err)
	}

	r := gin.Default()
	r.GET("/customer", customer.FindAllCustomerHandler)
	r.GET("/customer/:id", customer.FindCustomerHandler)
	r.POST("/customer", customer.CreateCustomerHandler)
	r.PUT("/customer/:id", customer.UpdateCustomerHandler)
	r.DELETE("/customer/:id")
	r.Run(":2009")
}

