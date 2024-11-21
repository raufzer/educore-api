package controllers

import (
	"educore-api/internal/models"
	"educore-api/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserController represents the controller for user operations
type UserController struct {
	UserService services.UserService
}

// NewUserController creates a new UserController
func NewUserController(userService services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user in the system
// @Tags user
// @Accept json
// @Produce json
// @Param user body models.User true "User data"
// @Success 201 {object} map[string]string "User created successfully"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /user/create [post]
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
	ctx.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

// GetUser godoc
// @Summary Get a user by username
// @Description Get a user from the system by their username
// @Tags user
// @Accept json
// @Produce json
// @Param name path string true "Username"
// @Success 200 {object} models.User "User found"
// @Failure 404 {object} map[string]string "User not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /user/get/{name} [get]
func (uc *UserController) GetUser(ctx *gin.Context) {
	username := ctx.Param("name")
	user, err := uc.UserService.GetUser(&username)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Get a list of all users
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {array} models.User "List of users"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /user/getall [get]
func (uc *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := uc.UserService.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// UpdateUser godoc
// @Summary Update an existing user
// @Description Update an existing user's details
// @Tags user
// @Accept json
// @Produce json
// @Param user body models.User true "User data"
// @Success 200 {object} map[string]string "User updated successfully"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /user/update [patch]
func (uc *UserController) UpdateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserService.UpdateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User updated"})
}

// DeleteUser godoc
// @Summary Delete a user by username
// @Description Delete a user from the system by their username
// @Tags user
// @Accept json
// @Produce json
// @Param name path string true "Username"
// @Success 200 {object} map[string]string "User deleted successfully"
// @Failure 404 {object} map[string]string "User not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /user/delete/{name} [delete]
func (uc *UserController) DeleteUser(ctx *gin.Context) {
	username := ctx.Param("name")
	err := uc.UserService.DeleteUser(&username)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
