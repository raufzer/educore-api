package main

import (
	"educore-api/config"
	"educore-api/internal/controllers"
	"educore-api/internal/repositories"
	"educore-api/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	appConfig := config.LoadConfig()

	// Set up the database
	dbConfig := config.ConnectDatabase(appConfig)
	defer dbConfig.Client.Disconnect(dbConfig.Ctx)

	// Initialize dependencies
	userRepo := repositories.NewUserRepository(dbConfig.UserCollection, dbConfig.Ctx)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	// Set up the Gin server
	server := gin.Default()
	basePath := server.Group("/v1")
	userController.RegisterUserRoutes(basePath)

	// Start the server
	log.Printf("Server running on port %s", appConfig.ServerPort)
	log.Fatal(server.Run(":" + appConfig.ServerPort))
}
