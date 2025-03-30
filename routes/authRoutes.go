package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

// SetupAuthRoutes registers the authentication routes
func SetupAuthRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/signup", controllers.Signup)
		auth.POST("/login", controllers.Login)
	}
}
