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
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, customer)
}

func UpdateCustomerHandler(c *gin.Context) {

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
