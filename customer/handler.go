package customer

import (
	"github.com/gin-gonic/gin"
	"github.com/masinew/gofinal/customer/data"
	"github.com/masinew/gofinal/database"
	"net/http"
	"strconv"
)

func FindAllCustomerHandler(c *gin.Context) {
	conn := database.Connection()
	customers := data.FindAll(conn)
	c.JSON(http.StatusOK, customers)
}

func FindCustomerHandler(c *gin.Context) {
	id, err := canParsingId(c)
	if err != nil {
		return
	}

	conn := database.Connection()
	customer := data.FindOne(conn, id)
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

	c.JSON(http.StatusOK, customer)
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

func isMissingId(c *gin.Context) string {
	param := c.Param("id")
	if param == "" {
		c.JSON(http.StatusBadRequest, data.Message{
			Message: "Missing id query param",
		})

		return ""
	}

	return param
}

func canParsingId(c *gin.Context) (int, error) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, data.Message{
			Message: "id query param is not an integer",
		})
		return 0, err
	}

	return id, nil
}

func canBindJson(c *gin.Context) (*data.Customer, error) {
	var customer *data.Customer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, data.Message{
			Message: "request body is invalid",
		})
		return customer, err
	}

	return customer, nil
}

func validateCustomer(c *gin.Context, customer data.Customer) bool {
	var valid = true
	if (customer.Name == "" && customer.Email == "") || customer.Status == "" {
		valid = false
	}

	if !valid {
		c.JSON(http.StatusBadRequest, data.Message{
			Message: "request body is invalid",
		})
	}

	return valid
}

