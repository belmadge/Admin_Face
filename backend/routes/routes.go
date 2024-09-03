package routes

import (
	"github.com/belmadge/Admin_Face/handlers"
	"github.com/belmadge/Admin_Face/utils"
	"github.com/gin-gonic/gin"
)

// InitializeRoutes initializes all routes of the application
func InitializeRoutes(router *gin.Engine) {
	// Authentication routes
	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)

	// Protected routes
	api := router.Group("/api")
	api.Use(utils.AuthMiddleware())
	{
		// User routes
		api.GET("/users", handlers.GetUsers)
		api.GET("/users/:id", handlers.GetUser)
		api.PUT("/users/:id", handlers.UpdateUser)
		api.DELETE("/users/:id", handlers.DeleteUser)

		// Event routes
		api.GET("/events", handlers.GetEvents)
		api.POST("/events", handlers.CreateEvent)
		api.PUT("/events/:id", handlers.UpdateEvent)
		api.DELETE("/events/:id", handlers.DeleteEvent)

		// Notification routes
		api.GET("/notifications", handlers.GetNotifications)
		api.POST("/notifications", handlers.CreateNotification)
		api.DELETE("/notifications/:id", handlers.DeleteNotification)

		// Export routes
		api.GET("/export/events", handlers.ExportEvents)
	}
}
