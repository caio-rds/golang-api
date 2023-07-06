package user

import (
	userModels "github.com/caio-rds/golang-api/src/model/user/requests"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (uc *Controller) Username(c *gin.Context) {
	var user userModels.FindByUsername
	if err := c.ShouldBindUri(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := uc.db.GetUserByName(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result == nil {
		c.JSON(http.StatusOK, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, result)

}

func (uc *Controller) Email(c *gin.Context) {
	var user userModels.FindUserByEmail
	if err := c.ShouldBindUri(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := uc.db.GetUserByEmail(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if result == nil {
		c.JSON(http.StatusOK, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, result)
}
