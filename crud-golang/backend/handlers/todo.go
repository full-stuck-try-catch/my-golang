package handlers

import (
	"net/http"
	"todo-backend/database"
	"todo-backend/models"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	var todos []models.Todo
	database.DB.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&todo)
	c.JSON(http.StatusOK, todo)
}

func UpdateTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")
	database.DB.First(&todo, id)

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&todo)
	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&models.Todo{}, id)
	c.Status(http.StatusOK)
}
