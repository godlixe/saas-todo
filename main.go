package main

import (
	"context"
	"fmt"
	"os"
	"saas-todo-list/internal/app/services"
	"saas-todo-list/internal/infrastructure/storage/postgresql"
	"saas-todo-list/internal/presentation/controllers"
	"saas-todo-list/internal/presentation/middlewares"
	"saas-todo-list/internal/presentation/routes"
	"saas-todo-list/pkg/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	db, err := database.NewPostgresClient(database.DatabaseCredentials{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	logger := zerolog.New(os.Stdout)

	userRepository := postgresql.NewUserRepository(context.Background(), db)
	todoRepository := postgresql.NewTodoRepository(context.Background(), db)

	authService := services.NewAuthService(userRepository)
	todoService := services.NewTodoService(todoRepository)

	authController := controllers.NewAuthController(authService)
	todoController := controllers.NewTodoController(todoService)

	server := gin.Default()
	server.Use(middlewares.ErrorHandler(logger))

	routes.AuthRoutes(server, authController)
	routes.TodoRoutes(server, todoController)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server.Run(":" + port)
}
