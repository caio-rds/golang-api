package user

import (
	"github.com/caio-rds/golang-api/src/configurations/rest_err"
	"github.com/caio-rds/golang-api/src/configurations/validation"
	"github.com/caio-rds/golang-api/src/database"
	userReq "github.com/caio-rds/golang-api/src/model/user/requests"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	db *database.Db
}

func NewController(db *database.Db) *Controller {
	return &Controller{db: db}
}

func (uc *Controller) CreateUser(c *gin.Context) {
	var userPayload userReq.Request

	if err := c.ShouldBindJSON(&userPayload); err != nil {
		restErr := validation.ValidateError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	result, err := uc.db.NewDbUser(userPayload)
	if err != nil {
		valError := rest_err.NewBadRequestError(err.Error())
		c.JSON(valError.Code, valError)
		return
	}

	c.JSON(http.StatusOK, result)

}
