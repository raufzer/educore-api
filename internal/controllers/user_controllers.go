package controllers

import (
	"educore-api/data/request"
	"educore-api/data/response"
	"educore-api/internal/models"
	"educore-api/internal/services"
	"educore-api/pkg/helpers"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	// Swagger imports
	_ "educore-api/docs" // This will be generated by swag
)

// @title           EduCore API
// @version         1.0
// @description     User Management API for EduCore Platform
// @host            localhost:8080
// @BasePath        /api/v1

// UserController godoc
// @Description Handles user-related operations
type UserController struct {
	UserService services.UserService
}

// NewUserController creates a new UserController
func NewUserController(userService services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

// toUserResponse converts a User model to UserResponse
func toUserResponse(user *models.User) response.UserResponse {
	return response.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// CreateUser godoc
// @Summary Create a new user
// @Description Creates a new user with the provided details
// @Tags Users
// @Accept json
// @Produce json
// @Param user body request.CreateUsersRequest true "User Creation Request"
// @Success 201 {object} response.Response{data=response.UserResponse} "User created successfully"
// @Failure 400 {object} response.Response "Bad request - invalid input"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /users [post]
func (uc *UserController) CreateUser(ctx *gin.Context) {
	var req request.CreateUsersRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:    http.StatusBadRequest,
			Status:  "BAD_REQUEST",
			Message: err.Error(),
		})
		return
	}

	user := models.User{
		ID:        primitive.NewObjectID(),
		Name:      req.Name,
		Email:     req.Email,
		Password:  req.Password,
		Role:      req.Role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := uc.UserService.CreateUser(&user)
	if err != nil {
		helpers.ErrorPanic(err)
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Code:    http.StatusInternalServerError,
			Status:  "INTERNAL_SERVER_ERROR",
			Message: "Failed to create user",
		})
		return
	}

	userResponse := toUserResponse(&user)
	ctx.JSON(http.StatusCreated, response.Response{
		Code:    http.StatusCreated,
		Status:  "CREATED",
		Message: "User created successfully",
		Data:    userResponse,
	})
}

// UpdateUser godoc
// @Summary Update an existing user
// @Description Updates user details by username
// @Tags Users
// @Accept json
// @Produce json
// @Param name path string true "Username"
// @Param user body request.UpdateUserRequest true "User Update Request"
// @Success 200 {object} response.Response{data=response.UserResponse} "User updated successfully"
// @Failure 400 {object} response.Response "Bad request - invalid input"
// @Failure 404 {object} response.Response "User not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /users/{name} [put]
func (uc *UserController) UpdateUser(ctx *gin.Context) {
	var req request.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:    http.StatusBadRequest,
			Status:  "BAD_REQUEST",
			Message: err.Error(),
		})
		return
	}

	username := ctx.Param("name")

	existingUser, err := uc.UserService.GetUser(&username)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.Response{
			Code:    http.StatusNotFound,
			Status:  "NOT_FOUND",
			Message: "User not found",
		})
		return
	}

	if req.Name != "" {
		existingUser.Name = req.Name
	}
	if req.Email != "" {
		existingUser.Email = req.Email
	}
	if req.Password != "" {
		existingUser.Password = req.Password
	}
	if req.Role != "" {
		existingUser.Role = req.Role
	}
	existingUser.UpdatedAt = time.Now()

	err = uc.UserService.UpdateUser(existingUser)
	if err != nil {
		helpers.ErrorPanic(err)
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Code:    http.StatusInternalServerError,
			Status:  "INTERNAL_SERVER_ERROR",
			Message: "Failed to update user",
		})
		return
	}

	userResponse := toUserResponse(existingUser)
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "User updated successfully",
		Data:    userResponse,
	})
}

// GetUser godoc
// @Summary Retrieve a user by username
// @Description Fetches user details for a specific username
// @Tags Users
// @Produce json
// @Param name path string true "Username"
// @Success 200 {object} response.Response{data=response.UserResponse} "User found"
// @Failure 404 {object} response.Response "User not found"
// @Router /users/{name} [get]
func (uc *UserController) GetUser(ctx *gin.Context) {
	username := ctx.Param("name")
	user, err := uc.UserService.GetUser(&username)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.Response{
			Code:    http.StatusNotFound,
			Status:  "NOT_FOUND",
			Message: "User not found",
		})
		return
	}

	userResponse := toUserResponse(user)
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "User found",
		Data:    userResponse,
	})
}

// GetAllUsers godoc
// @Summary List all users
// @Description Retrieves a list of all users in the system
// @Tags Users
// @Produce json
// @Success 200 {object} response.Response{data=[]response.UserResponse} "Users retrieved successfully"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /users [get]
func (uc *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := uc.UserService.GetAllUsers()
	if err != nil {
		helpers.ErrorPanic(err)
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Code:    http.StatusInternalServerError,
			Status:  "INTERNAL_SERVER_ERROR",
			Message: "Failed to fetch users",
		})
		return
	}

	userResponses := make([]response.UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = toUserResponse(user)
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Users retrieved successfully",
		Data:    userResponses,
	})
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Deletes a user by username
// @Tags Users
// @Param name path string true "Username"
// @Success 200 {object} response.Response "User deleted successfully"
// @Failure 404 {object} response.Response "User not found"
// @Router /users/{name} [delete]
func (uc *UserController) DeleteUser(ctx *gin.Context) {
	username := ctx.Param("name")
	err := uc.UserService.DeleteUser(&username)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.Response{
			Code:    http.StatusNotFound,
			Status:  "NOT_FOUND",
			Message: "User not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "User deleted successfully",
	})
}
