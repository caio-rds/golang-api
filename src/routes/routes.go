package routes

import (
	"github.com/caio-rds/golang-api/src/controller/user"
	"github.com/caio-rds/golang-api/src/database"
	"github.com/gin-gonic/gin"
)

func InitRoutes(db *database.Db) error {
	userController := user.NewController(db)
	r := gin.Default()
	userGroup := r.Group("/user")
	{
		userGroup.GET("/:username", userController.Username)
		userGroup.GET("/by_email/:email", userController.Email)
		userGroup.POST("/", userController.CreateUser)
		userGroup.PUT("/", userController.UpdateUser)
		userGroup.DELETE("/:username", userController.DeleteUser)
	}
	if err := r.Run(); err != nil {
		return err
	}
	return nil
}
