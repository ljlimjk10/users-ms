package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	users_controllers "github.com/ljlimjk10/users-ms/controllers"
	"github.com/ljlimjk10/users-ms/db"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env: %s", err)
	}
	db := db.InitDBConnection()
	router := gin.Default()
	users := router.Group("/users")
	{
		users.GET("/info", users_controllers.TestController)
		users.POST("/register", func(c *gin.Context) {
			users_controllers.RegisterUser(c, db.DB)
		})
	}

	router.Run("localhost:8083")

}
