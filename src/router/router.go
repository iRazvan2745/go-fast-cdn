package router

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/kevinanielsen/go-fast-cdn/src/middleware"
)

// Router initializes the router and sets up middleware, routes, etc.
// It returns a *gin.Engine instance configured with the routes, middleware, etc.
func Router() {
	port := ":" + os.Getenv("PORT")

	s := NewServer(
		WithPort(port),
		WithMiddleware(cors.New(cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			AllowWebSockets:  true,
			MaxAge:           86400,
		})),
	)

	// Add all the API routes
	s.AddApiRoutes()

	s.Run()
}
