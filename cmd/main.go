package main

import (
	"educore-api/config"
	"educore-api/internal/controllers"
	"educore-api/internal/repositories"
	"educore-api/internal/services"
	v1 "educore-api/routes/api/v1"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// Swagger documentation
	_ "educore-api/docs"
)

// @title           EduCore API
// @version         1.0
// @description     Complete API for EduCore Platform
// @host            localhost:8080
// @BasePath        /v1
func main() {
	// Load configuration
	appConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Database connection
	dbConfig := config.ConnectDatabase(appConfig)
	defer func() {
		if err := dbConfig.Client.Disconnect(dbConfig.Ctx); err != nil {
			log.Printf("Error disconnecting database: %v", err)
		}
	}()

	// Validator
	validate := validator.New()

	// Repositories
	userRepo := repositories.NewUserRepository(dbConfig.UserCollection)

	// Services
	userService := services.NewUserService(userRepo)
	authService := services.NewAuthenticationServiceImpl(userRepo, validate)

	// Controllers
	userController := controllers.NewUserController(userService)
	authController := controllers.NewAuthenticationController(authService)

	// Gin setup
	gin.SetMode(gin.ReleaseMode) // Production mode
	server := gin.Default()

	// CORS Configuration
	server.Use(cors.Default())

	// Global middleware
	server.Use(gin.Recovery())

	// Swagger setup
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("http://localhost:"+appConfig.ServerPort+"/swagger/doc.json"),
	))

	// API Routes
	basePath := server.Group("/v1")
	v1.UserRoutes(basePath, userController)
	v1.AuthenticationRoutes(basePath, authController)

	// Server startup
	serverAddr := ":" + appConfig.ServerPort
	log.Printf("🚀 Server starting on %s", serverAddr)
	log.Fatal(server.Run(serverAddr))
}
