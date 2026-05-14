package routes

import (
	"hello_project/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", handlers.Home)
	r.GET("/ping", handlers.Ping)
	r.GET("/comments", handlers.GetComments)

	users := r.Group("/users")
	{
		users.GET("", handlers.GetUsers)
		users.GET("/:id", handlers.GetUserByID)
		users.POST("", handlers.CreateUser)
		users.PUT("/:id", handlers.UpdateUser)
		users.DELETE("/:id", handlers.DeleteUser)
	}

	r.POST("/seed/users", handlers.SeedUsers)
}
