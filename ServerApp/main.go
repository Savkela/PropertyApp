package main

import (
	"ServerApp/controller"
	"ServerApp/initializers"
	"ServerApp/service"

	"github.com/gin-gonic/gin"
)

var (
	userService    service.UserService       = service.NewUserService()
	userController controller.UserController = controller.NewUserController(userService)
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	server := gin.New()

	server.Use(gin.Recovery(), gin.Logger(), gin.BasicAuth(gin.Accounts{"savic": "nikola"}))

	//users
	server.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, userController.FindAll())
	})

	server.POST("/users", userController.Save)

	server.GET("/users/:id", userController.FindOne)

	server.PUT("/users/:id", userController.Update)

	server.DELETE("/users/:id", userController.Delete)

	server.Run()

}
