package api

import (
	"net/http"

	"github.com/clerk/clerk-sdk-go/v2"
	clerkHttp "github.com/clerk/clerk-sdk-go/v2/http"
	"github.com/clerk/clerk-sdk-go/v2/user"
	"github.com/gin-gonic/gin"
)

func WithClerkAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c.Request = r
			c.Next()
		})

		if c.Writer.Status() == http.StatusUnauthorized {
			c.JSON(http.StatusUnauthorized, ErrorResponse{
				Error: "Unauthorized",
			})

			c.Abort()
			return
		}

		clerkHttp.WithHeaderAuthorization()(handler).ServeHTTP(c.Writer, c.Request)
	}
}

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, ok := clerk.SessionClaimsFromContext(c.Request.Context())
		if !ok {
			c.JSON(http.StatusUnauthorized, ErrorResponse{
				Error: "Unauthorized",
			})

			c.Abort()
			return
		}

		user, err := user.Get(c.Request.Context(), claims.Subject)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Error: "Failed to find user!",
			})

			c.Abort()
			return
		}

		c.Set("userId", user.ID)
		c.Next()
	}
}
