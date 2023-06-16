package user

import (
	"github.com/caio-rds/golang-api/src/database"
	userModels "github.com/caio-rds/golang-api/src/model/user/requests"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UsernameFind struct {
	db *database.Db
}

func NewUsernameFind(db *database.Db) *UsernameFind {
	return &UsernameFind{db: db}
}

func (uf UsernameFind) FindByUsername(c *gin.Context) {
	var user userModels.FindByUsernameRequest
	if err := c.ShouldBindUri(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := uf.db.GetUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)

	//c.JSON(http.StatusOK, gin.H{"username": user.Username})
}

func FindUserByEmail(c *gin.Context) {
	var user userModels.FindUserByEmailRequest
	if err := c.ShouldBindUri(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"email": user.Email})
}
