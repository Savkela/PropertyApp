package controller

import (
	"ServerApp/entity"
	"ServerApp/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Save(ctx *gin.Context)
	FindAll() []entity.User
	FindOne(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type userController struct {
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &userController{
		service: service,
	}
}

func (c *userController) FindAll() []entity.User {
	return c.service.FindAll()
}

func (c *userController) FindOne(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := c.service.FindOne(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *userController) Save(ctx *gin.Context) {
	var user entity.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	err := c.service.Save(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *userController) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	var user entity.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	updatedUser, err := c.service.Update(id, user)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "The user is not found"})
		return
	}

	ctx.JSON(http.StatusOK, updatedUser)
}

func (c *userController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found or error deleting"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "The user was successfully deleted"})
}
