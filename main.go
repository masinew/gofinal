package main

import (
	"github.com/gin-gonic/gin"
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
	r.GET("/customer", handler)
	r.Run(":2009")
}

func handler(c *gin.Context) {

}
