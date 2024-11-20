package main

import (
	"educore-api/config"
	"educore-api/internal/controllers"
	"educore-api/internal/repositories"
	"educore-api/internal/services"
	"educore-api/routes/api/v1"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	appConfig := config.LoadConfig()

	dbConfig := config.ConnectDatabase(appConfig)
	defer dbConfig.Client.Disconnect(dbConfig.Ctx)

	userRepo := repositories.NewUserRepository(dbConfig.UserCollection)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	server := gin.Default()
	basePath := server.Group("/v1")

	v1.RegisterUserRoutes(basePath, userController)

	log.Printf("Server running on port %s", appConfig.ServerPort)
	log.Fatal(server.Run(":" + appConfig.ServerPort))
}
