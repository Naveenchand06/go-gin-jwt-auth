package routes

import (
	"github.com/Naveenchand06/go-gin-jwt-auth/controllers"
	"github.com/gin-gonic/gin"
)


func AuthRoutes(router *gin.Engine) {
	// * Grouping `/auth` routes
	authRoutesGroup := router.Group("/auth")

	// * Group Routes
	authRoutesGroup.POST("/auth/login", controllers.Login)
	authRoutesGroup.POST("/auth/signup", controllers.Signup)
}