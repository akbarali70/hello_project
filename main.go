package main

import (
	"log"

	_ "hello_project/docs"
	"hello_project/internal/db"
	"hello_project/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title Hello Project API
// @version 1.0
// @description Gin + PostgreSQL CRUD API
// @host localhost:8001
// @BasePath /
func main() {
	_ = godotenv.Load()

	gin.SetMode(gin.ReleaseMode)

	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.SetTrustedProxies(nil)

	routes.RegisterRoutes(r)

	r.Run("0.0.0.0:8001")
}
