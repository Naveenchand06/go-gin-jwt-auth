package routes

import (
	"github.com/Naveenchand06/go-gin-jwt-auth/controllers"
	"github.com/Naveenchand06/go-gin-jwt-auth/middleware"
	"github.com/gin-gonic/gin"
)


func UserRoutes(router *gin.Engine) {
	// * Grouping `/users` routes
	userRoutesGroup := router.Group("/users")
	// * Adding middleware for authorization
	userRoutesGroup.Use(middleware.Authenticate)

	// * Group Routes
	userRoutesGroup.GET("/", controllers.GetUsers)
	userRoutesGroup.GET("/:id", controllers.GetUserById)
}