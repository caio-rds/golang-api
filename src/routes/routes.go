package routes

import (
	"github.com/caio-rds/golang-api/src/controller/user"
	"github.com/caio-rds/golang-api/src/database"
	"github.com/gin-gonic/gin"
)

func InitRoutes(db *database.Db) error {
	userController := user.NewController(db)
	userFinder := user.NewUsernameFind(db)
	r := gin.Default()
	userGroup := r.Group("/user")
	{
		userGroup.GET("/:username", userFinder.FindByUsername)
		userGroup.GET("/by_email/:email", user.FindUserByEmail)
		userGroup.POST("/", userController.CreateUser)
		userGroup.PUT("/:id")
		userGroup.DELETE("/:id")
	}
	if err := r.Run(); err != nil {
		return err
	}
	return nil
}
