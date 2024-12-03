package v1

import (
	"educore-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func AuthenticationRoutes(rg *gin.RouterGroup, authenticationController *controllers.AuthenticationController) {
	rg.POST("/sessions", authenticationController.Login)
	rg.POST("/register", authenticationController.Register)
}
