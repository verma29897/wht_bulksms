package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/verma29897/bulksms/handlers"
	"github.com/verma29897/bulksms/middleware"
)

func RegisterRoutes(router *gin.Engine) {
	// Health Check
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "WhatsApp Messaging API is live"})
	})

	// Public auth endpoints
	router.POST("/auth/signup", handlers.Signup)
	router.POST("/auth/login", handlers.Login)

	// Embedded Signup endpoints (public)
	router.GET("/onboard/callback", handlers.OnboardingCallback)
	router.POST("/store-onboarding", handlers.StoreOnboarding)

	// Protected API
	api := router.Group("")
	api.Use(middleware.AuthMiddleware())

	// Template Management
	api.GET("/templates/:waba_id", handlers.FetchTemplates) // Fetch templates
	api.POST("/templates", handlers.CreateTemplate)         // Create template

	// Media Upload
	api.POST("/upload", handlers.UploadHeaderHandle) // Get media_id by uploading file
	api.POST("/upload/header", handlers.UploadMedia) // Upload header media (handle_header)

	// Messaging
	api.POST("/send", handlers.SendMessagesHandler) // Send message using template

	// (removed) Store onboarding is public so success page can call it without JWT

	// GORM backed sample endpoints
	api.GET("/users", handlers.ListUsers)
	api.GET("/accounts", handlers.ListAccounts)
	api.GET("/me", handlers.Me)
}
