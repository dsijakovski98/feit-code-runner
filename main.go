package main

import (
	"fmt"
	"log"
	"net/http"

	"os"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/dsijakovski98/feit-code-runner/api"
	"github.com/dsijakovski98/feit-code-runner/config"
	"github.com/dsijakovski98/feit-code-runner/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	if err := os.Mkdir(utils.TMP_RUN_FILE, os.ModePerm); err != nil {
		log.Printf("Error creating temp run file: %v", err)
	}

	if err := os.Mkdir(utils.TMP_TESTS_FILE, os.ModePerm); err != nil {
		log.Printf("Error creating temp tests file: %v", err)
	}
}

func main() {
	// Get the secret key from environment variable
	secretKey := os.Getenv("CLERK_SECRET")
	if secretKey == "" {
		log.Fatal("CLERK_SECRET not set in environment")
	}

	clerk.SetKey(secretKey)

	router := gin.Default()

	fmt.Println(config.GetAllowedOrigins())

	// Configure CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     config.GetAllowedOrigins(),
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "FEIT Code Runner is locked in and ready to go!")
	})

	router.Use(api.WithClerkAuth(), api.Authorize())

	router.POST("/run", api.RunCodeHandler)
	router.POST("/cleanup", api.CleanupDebugHandler)

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	if err := router.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
