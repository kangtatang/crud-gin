package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inisialisasi database
	db := InitDB()

	// Inisialisasi Gin
	r := gin.Default()

	// Inisialisasi handler dengan database
	userHandler := NewUserHandler(db)

	// Routing user management
	r.GET("/users", userHandler.GetUsers)
	r.GET("/users/:id", userHandler.GetUserByID)
	r.POST("/users", userHandler.CreateUser)
	r.PUT("/users/:id", userHandler.UpdateUser)
	r.DELETE("/users/:id", userHandler.DeleteUser)

	// Menjalankan server
	log.Fatal(r.Run(":8080"))
}
