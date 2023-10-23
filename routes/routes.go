package routes

import (
	"todogorest/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter(todoController *controllers.TodoController) *gin.Engine {
	router := gin.Default()

	baseRouter := router.Group("/api")

	// TODOS
	todoRouter := baseRouter.Group("/todos")

	todoRouter.GET("/", todoController.GetAllTodos)

	todoRouter.POST("/", todoController.CreateTodo)

	todoRouter.DELETE("/:todo_id", todoController.DeleteTodo)

	todoRouter.GET("/:todo_id", todoController.GetTodo)

	todoRouter.PUT("/:todo_id", todoController.UpdateTodo)

	return router
}
