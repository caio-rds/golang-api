package user

import (
	userModels "github.com/caio-rds/golang-api/src/model/requests/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindUserById(c *gin.Context) {
	var user userModels.FindUserByIdRequest
	if err := c.ShouldBindUri(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": user.ID})
}

func FindUserByEmail(c *gin.Context) {

}
