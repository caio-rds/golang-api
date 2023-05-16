package routes

import (
	"github.com/caio-rds/golang-api/src/controller/user"
	"github.com/gin-gonic/gin"
)

func InitRoutes() error {
	r := gin.Default()
	userGroup := r.Group("/user")
	{
		userGroup.GET("/by_id/:id", user.FindUserById)
		userGroup.GET("/by_email/:email", user.FindUserByEmail)
		userGroup.POST("/")
		userGroup.PUT("/:id")
		userGroup.DELETE("/:id")
	}
	if err := r.Run(); err != nil {
		return err
	}
	return nil
}
