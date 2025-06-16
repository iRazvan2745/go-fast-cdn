package router

import (
	"os"

	"github.com/kevinanielsen/go-fast-cdn/src/middleware"
)

// Router initializes the router and sets up middleware, routes, etc.
// It returns a *gin.Engine instance configured with the routes, middleware, etc.
func Router() {
	port := ":" + os.Getenv("PORT")

	s := NewServer(
		WithPort(port),
		WithMiddleware(middleware.CORSMiddleware()),
	)

	// Add all the API routes
	s.AddApiRoutes()

	s.Run()
}
