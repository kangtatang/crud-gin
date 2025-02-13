package main

import (
	"log"
	"os"

	"gin-user-management/database"
	"gin-user-management/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load konfigurasi dari .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Inisialisasi database
	database.InitDB()
	r := gin.Default()

	userHandler := handlers.NewUserHandler() // Gunakan handlers.NewUserHandler

	r.GET("/users", userHandler.GetUsers)
	r.GET("/users/:id", userHandler.GetUserByID)
	r.POST("/users", userHandler.CreateUser)
	r.PUT("/users/:id", userHandler.UpdateUser)
	r.DELETE("/users/:id", userHandler.DeleteUser)

	log.Fatal(r.Run(os.Getenv("SERVER_PORT")))
}
