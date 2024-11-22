package controllers

import (
	"educore-api/data/request"
	"educore-api/data/response"
	"educore-api/internal/services"
	"educore-api/pkg/helpers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthenticationController struct {
	AuthenticationService services.AuthenticationService
}

func NewAuthenticationController(service services.AuthenticationService) *AuthenticationController {
	return &AuthenticationController{AuthenticationService: service}
}

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
	// ctx.SetCookie(
	// 	"token",     // Cookie name
	// 	token,       // Token value
	// 	60*60,       // Max-Age in seconds (convert minutes to seconds)
	// 	"/",         // Path
	// 	"localhost", // Domain
	// 	false,       // Secure (set to true for HTTPS)
	// 	true,        // HttpOnly (prevents JavaScript access)
	// )

	ctx.JSON(http.StatusOK, webResponse)
}

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
