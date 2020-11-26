package customer

import (
	"github.com/gin-gonic/gin"
	"github.com/masinew/gofinal/customer/data"
	"net/http"
	"strconv"
)

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

