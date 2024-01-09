package main

import (
	"github.com/gin-gonic/gin"
	"github.com/your-module/controllers"
	"github.com/your-module/middleware"
)

func main() {
	r := gin.Default()

	r.POST("/login", controllers.Login)

	protected := r.Group("/api")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		protected.GET("/protected", controllers.Protected)
	}

	r.Run(":8080")
}
