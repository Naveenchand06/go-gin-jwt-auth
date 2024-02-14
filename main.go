package main

import (
	"fmt"
	"os"

	"github.com/Naveenchand06/go-gin-jwt-auth/constants"
	"github.com/Naveenchand06/go-gin-jwt-auth/database"
	"github.com/Naveenchand06/go-gin-jwt-auth/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func main() {
	// * Loading .env file
	envErr := godotenv.Load()

	// * If there is any error in loading .env file printing the error and exiting the main function
    if envErr != nil {
        fmt.Println("Error loading .env file -> ", envErr)
        return
    }

	// * Connecting to DB
	database.GetDB()

	// * Reading port number from .env file
	port := os.Getenv(constants.PORT)
	if port == "" {
		port = "5010"
	}

	// * Creatinga new gin router (*gin.Engine)
	router := gin.New()

	// * Adding middleware for logs
	router.Use(gin.Logger())

	// * Registering Routes
	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	// * Listening for requests
	router.Run(":" + port)
}