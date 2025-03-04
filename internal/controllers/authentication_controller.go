package controllers

import (
	"educore-api/data/request"
	"educore-api/data/response"
	"educore-api/internal/services"
	"educore-api/pkg/helpers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	// Swagger imports
	_ "educore-api/docs" // This will be generated by swag
)

// @title           EduCore Authentication API
// @version         1.0
// @description     Authentication Management API for EduCore Platform
// @host            localhost:8080
// @BasePath        /api/v1

// AuthenticationController godoc
// @Description Handles authentication-related operations
type AuthenticationController struct {
	AuthenticationService services.AuthenticationService
}

// NewAuthenticationController creates a new AuthenticationController
func NewAuthenticationController(service services.AuthenticationService) *AuthenticationController {
	return &AuthenticationController{AuthenticationService: service}
}

// Login godoc
// @Summary User Login
// @Description Authenticate user and generate access token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param loginRequest body request.LoginRequest true "Login Credentials"
// @Success 200 {object} response.Response{data=response.LoginResponse} "Successfully logged in"
// @Failure 400 {object} response.Response "Invalid username or password"
// @Router /auth/login [post]
func (controller *AuthenticationController) Login(ctx *gin.Context) {
	loginRequest := request.LoginRequest{}
	err := ctx.ShouldBindJSON(&loginRequest)
	helpers.ErrorPanic(err)

	token, err_token := controller.AuthenticationService.Login(loginRequest)
	fmt.Println(err_token)
	if err_token != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Invalid username or password",
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	resp := response.LoginResponse{
		TokenType: "Bearer",
		Token:     token,
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully log in!",
		Data:    resp,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

// Register godoc
// @Summary User Registration
// @Description Register a new user in the system
// @Tags Authentication
// @Accept json
// @Produce json
// @Param userRequest body request.CreateUsersRequest true "User Registration Details"
// @Success 200 {object} response.Response "Successfully created user"
// @Failure 400 {object} response.Response "Invalid registration details"
// @Router /auth/register [post]
func (controller *AuthenticationController) Register(ctx *gin.Context) {
	createUsersRequest := request.CreateUsersRequest{}
	err := ctx.ShouldBindJSON(&createUsersRequest)
	helpers.ErrorPanic(err)

	controller.AuthenticationService.Register(createUsersRequest)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully created user!",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
