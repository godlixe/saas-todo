package routes

import (
	"saas-todo-list/internal/presentation/controllers"
	"saas-todo-list/internal/presentation/middlewares"

	"github.com/gin-gonic/gin"
)

func TodoRoutes(router *gin.Engine, todoController *controllers.TodoController) {
	todoRoutes := router.Group("/todo", middlewares.Authenticate())
	{
		todoRoutes.GET("", todoController.GetAll)
		todoRoutes.GET(":id", todoController.GetById)
		todoRoutes.POST("", todoController.CreateTodo)
		todoRoutes.POST(":id", todoController.UpdateTodo)
		todoRoutes.DELETE(":id", todoController.Delete)
	}
}
