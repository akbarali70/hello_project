package routes

import (
	"hello_project/internal/handlers"
	"hello_project/internal/middleware"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", handlers.Home)
	r.GET("/ping", handlers.Ping)
	r.GET("/comments", handlers.GetComments)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	users := r.Group("/users")
	{
		users.GET("", middleware.OneRequestPer10Seconds(), handlers.GetUsers)
		users.GET("/:id", handlers.GetUserByID)
		users.POST("", handlers.CreateUser)
		users.PUT("/:id", handlers.UpdateUser)
		users.DELETE("/:id", handlers.DeleteUser)
	}

	r.POST("/seed/users", handlers.SeedUsers)
}
