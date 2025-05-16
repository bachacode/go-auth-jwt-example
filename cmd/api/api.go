package main

import (
	"fmt"

	"github.com/bachacode/go-auth-jwt-example/internal/database"
	"github.com/bachacode/go-auth-jwt-example/internal/handlers"
	"github.com/bachacode/go-auth-jwt-example/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	database.Init()
	database.Migrate()
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/signup", handlers.SignupHandler)
	router.POST("/login", handlers.LoginHandler)

	router.Use(middleware.AuthMiddleware()).GET("/validate", handlers.ValidateHandler)

	router.Run() // listen and serve on 0.0.0.0:8080
}
