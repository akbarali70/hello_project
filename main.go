package main

import (
	"log"

	"hello_project/internal/db"
	"hello_project/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.SetTrustedProxies(nil)

	routes.RegisterRoutes(r)

	r.Run(":8001")
}
