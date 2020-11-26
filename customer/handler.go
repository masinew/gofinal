package customer

import (
	"github.com/gin-gonic/gin"
	"github.com/masinew/gofinal/customer/data"
	"github.com/masinew/gofinal/database"
	"net/http"
)

func FindAllCustomerHandler(c *gin.Context) {
	conn := database.Connection()
	customers, err := data.FindAll(conn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, data.Message{
			Message: "The system error. Please contact admin.",
		})

		return
	}

	c.JSON(http.StatusOK, customers)
}

func FindCustomerHandler(c *gin.Context) {
	id, err := canParsingId(c)
	if err != nil {
		return
	}

	conn := database.Connection()
	customer, fErr := data.FindOne(conn, id)
	if fErr != nil {
		c.JSON(http.StatusInternalServerError, data.Message{
			Message: "The system error. Please contact admin.",
		})

		return
	}

	c.JSON(http.StatusOK, customer)
}

func CreateCustomerHandler(c *gin.Context) {
	customer, err := canBindJson(c)
	if err != nil || !validateCustomer(c, *customer){
		return
	}

	conn := database.Connection()
	iErr := data.Insert(conn, customer)
	if iErr != nil {
		c.JSON(http.StatusInternalServerError, data.Message{
			Message: "The system can not performs your action. Please contact admin.",
		})

		return
	}

	c.JSON(http.StatusCreated, customer)
}

func UpdateCustomerHandler(c *gin.Context) {
	_ = isMissingId(c)
	id, err := canParsingId(c)
	if err != nil {
		return
	}

	customer, bErr := canBindJson(c)
	if bErr != nil || !validateCustomer(c, *customer) {
		return
	}

	conn := database.Connection()
	uErr := data.Update(conn, id, customer)
	if uErr != nil {
		c.JSON(http.StatusInternalServerError, data.Message{
			Message: "The system can not performs your action. Please contact admin.",
		})

		return
	}

	c.JSON(http.StatusOK, customer)
}

func DeleteCustomerHandler(c *gin.Context) {
	_ = isMissingId(c)
	id, err := canParsingId(c)
	if err != nil {
		return
	}

	conn := database.Connection()
	uErr := data.Delete(conn, id)
	if uErr != nil {
		c.JSON(http.StatusInternalServerError, data.Message{
			Message: "The system can not performs your action. Please contact admin.",
		})

		return
	}

	c.JSON(http.StatusOK, data.Message{
		Message: "customer deleted",
	})
}

