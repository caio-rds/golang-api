package user

import (
	"fmt"
	"github.com/caio-rds/golang-api/src/configurations/validation"
	"github.com/caio-rds/golang-api/src/model/requests/user"
	userRes "github.com/caio-rds/golang-api/src/model/response/user"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userPayload user.Request

	if err := c.ShouldBindJSON(&userPayload); err != nil {
		restErr := validation.ValidateError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	var response = userRes.Response{
		ID:    "1",
		Email: userPayload.Email,
		Name:  userPayload.Name,
		Age:   userPayload.Age,
	}

	fmt.Println(response)
	c.JSON(200, &response)

}
