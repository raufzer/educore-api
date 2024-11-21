package main

import (
	"educore-api/config"
	"educore-api/internal/controllers"
	"educore-api/internal/repositories"
	"educore-api/internal/services"
	v1 "educore-api/routes/api/v1"
	"log"

	"github.com/gin-gonic/gin" // swagger embed files
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	_ "educore-api/docs" // Import the generated docs package
)

// @title EduCore API
// @version 1.0
// @description API for managing users in EduCore system
// @termsOfService http://example.com/terms/

// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:9090
// @BasePath /v1

func main() {
	appConfig, _ := config.LoadConfig()

	dbConfig := config.ConnectDatabase(appConfig)
	defer dbConfig.Client.Disconnect(dbConfig.Ctx)
	validate := validator.New()

	userRepo := repositories.NewUserRepository(dbConfig.UserCollection)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	authService := services.NewAuthenticationServiceImpl(userRepo, validate)
	authController := controllers.NewAuthenticationController(authService)

	server := gin.Default()

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	basePath := server.Group("/v1")
	v1.UserRoutes(basePath, userController)
	v1.AuthenticationRoutes(basePath, authController)

	log.Printf("Server running on port %s", appConfig.ServerPort)
	log.Fatal(server.Run(":" + appConfig.ServerPort))
}
