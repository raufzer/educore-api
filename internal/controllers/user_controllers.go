package controllers

import (
	"educore-api/internal/models"
	"educore-api/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User created"})
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	
	ctx.JSON(200, gin.H{"message": "User fetched"})
}

func (uc *UserController) GetAllUsers(ctx *gin.Context) {

	ctx.JSON(200, gin.H{"message": "All users fetched"})
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {

	ctx.JSON(200, gin.H{"message": "User updated"})
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {

	ctx.JSON(200, gin.H{"message": "User deleted"})
}
