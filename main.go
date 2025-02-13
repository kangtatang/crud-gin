package main

import (
	"log"

	"gin-user-management/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	db := InitDB()
	r := gin.Default()

	userHandler := handlers.NewUserHandler(db) // Gunakan handlers.NewUserHandler

	r.GET("/users", userHandler.GetUsers)
	r.GET("/users/:id", userHandler.GetUserByID)
	r.POST("/users", userHandler.CreateUser)
	r.PUT("/users/:id", userHandler.UpdateUser)
	r.DELETE("/users/:id", userHandler.DeleteUser)

	log.Fatal(r.Run(":8080"))
}
