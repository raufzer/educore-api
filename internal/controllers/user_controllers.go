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
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

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
