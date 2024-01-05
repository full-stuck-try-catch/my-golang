package main

import (
	"todo-backend/database"
	"todo-backend/handlers"
	"todo-backend/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.Connect()
	database.DB.AutoMigrate(&models.Todo{})

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Next()
	})

	r.GET("/api/todos", handlers.GetTodos)
	r.POST("/api/todos", handlers.CreateTodo)
	r.PUT("/api/todos/:id", handlers.UpdateTodo)
	r.DELETE("/api/todos/:id", handlers.DeleteTodo)

	r.Run(":8080")
}
