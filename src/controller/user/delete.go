package user

import (
	"github.com/caio-rds/golang-api/src/configurations/rest_err"
	user "github.com/caio-rds/golang-api/src/model/user/requests"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (uc *Controller) DeleteUser(c *gin.Context) {
	var userModel user.FindByUsername

	if err := c.ShouldBindUri(&userModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userDel, err := uc.db.DeleteUser(userModel)
	if err != nil {
		valError := rest_err.NewBadRequestError(err.Error())
		c.JSON(valError.Code, valError)
		return
	}
	if userDel == nil {
		c.JSON(http.StatusOK, gin.H{"error": "user not exists"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted": userDel.ID})
	return
}
