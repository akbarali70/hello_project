package main

import (
	"log"
	"os"

	_ "hello_project/docs"
	"hello_project/internal/db"
	"hello_project/internal/models"
	"hello_project/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	gin.SetMode(gin.ReleaseMode)

	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}

	db.DB.AutoMigrate(&models.User{})

	port := os.Getenv("APP_PORT")

	if port == "" {
		port = "8000"
	}

	r := gin.Default()
	r.SetTrustedProxies(nil)

	routes.RegisterRoutes(r)

	r.Run("0.0.0.0:" + port)
}
