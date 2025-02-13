package main

import (
	"log"

	"gin-user-management/database"
	"gin-user-management/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// db := database.InitDB()
	database.InitDB()
	r := gin.Default()

	userHandler := handlers.NewUserHandler() // Gunakan handlers.NewUserHandler

	r.GET("/users", userHandler.GetUsers)
	r.GET("/users/:id", userHandler.GetUserByID)
	r.POST("/users", userHandler.CreateUser)
	r.PUT("/users/:id", userHandler.UpdateUser)
	r.DELETE("/users/:id", userHandler.DeleteUser)

	log.Fatal(r.Run(":8080"))
}
