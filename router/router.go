package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shehbaazsk/go-commerce/middlewares"
)

// SetupRouter creates the main Gin engine and mounts all app routes
func SetupRouter(db *pgxpool.Pool) *gin.Engine {
	r := gin.New()

	// Attach middleware
	r.Use(gin.Recovery())
	r.Use(middlewares.CustomLogger())

	// Public routes group (no auth)
	public := r.Group("/api/v1")
	{
		// Health check route
		public.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "API is running",
			})
		})

		// Register public routes from apps here
	}

	// Protected routes group (JWT Auth + RBAC middleware)
	protected := r.Group("/api/v1")
	protected.Use(middlewares.JWTAuthMiddleware())
	{
		// IRegister protected routes from apps here
	}

	return r
}
