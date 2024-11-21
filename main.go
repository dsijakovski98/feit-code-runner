package main

import (
	"log"
	"net/http"
	"os"

	"github.com/clerk/clerk-sdk-go/v2"
	clerkHttp "github.com/clerk/clerk-sdk-go/v2/http"
	"github.com/clerk/clerk-sdk-go/v2/user"
	"github.com/dsijakovski98/feit-code-runner/api"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	if err := os.Mkdir("_tmp", os.ModePerm); err != nil {
		log.Printf("Error creating _tmp file: %v", err)
	}
}

func main() {
	// Get the secret key from environment variable
	secretKey := os.Getenv("CLERK_SECRET_KEY")
	if secretKey == "" {
		log.Fatal("CLERK_SECRET_KEY not set in environment")
	}

	clerk.SetKey(secretKey)

	router := gin.Default()

	// Configure CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3001"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.Use(withClerkAuth(), authorize())
	router.POST("/run", api.RunCodeHandler)

	if err := router.Run(":4000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func withClerkAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c.Request = r
			c.Next()
		})

		clerkHttp.WithHeaderAuthorization()(handler).ServeHTTP(c.Writer, c.Request)

		if c.Writer.Status() == http.StatusUnauthorized {
			c.JSON(http.StatusUnauthorized, api.ErrorResponse{
				Error: "Unauthorized",
			})

			c.Abort()
		}
	}
}

func authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, ok := clerk.SessionClaimsFromContext(c.Request.Context())
		if !ok {
			c.JSON(http.StatusUnauthorized, api.ErrorResponse{
				Error: "Unauthorized",
			})

			c.Abort()
			return
		}

		user, err := user.Get(c.Request.Context(), claims.Subject)
		if err != nil {
			c.JSON(http.StatusInternalServerError, api.ErrorResponse{
				Error: "Failed to find user!",
			})

			c.Abort()
			return
		}

		c.Set("userId", user.ID)
		c.Next()
	}
}
