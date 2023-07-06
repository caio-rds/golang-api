package user

import (
	userModels "github.com/caio-rds/golang-api/src/model/user/requests"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (uc *Controller) UpdateUser(c *gin.Context) {
	var editableUser userModels.EditUser

	if err := c.ShouldBindJSON(&editableUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username := userModels.FindByUsername{Username: editableUser.Username}

	userExists, err := uc.db.GetUserByName(username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if userExists == nil {
		c.JSON(http.StatusOK, gin.H{"error": "User not found"})
		return
	}
	updated, err := uc.db.EditUser(editableUser, username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updated)
	return
}
