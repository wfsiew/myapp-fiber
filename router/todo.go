package router

import (
    "app/controller"
    "github.com/gofiber/fiber/v2"
    "app/middleware"
)

func SetupTodoRoutes(router fiber.Router) {
    todo := router.Group("/todo")
    todo.Use(middleware.JWTProtected)
    todo.Post("/", controller.CreateTodo)
    todo.Get("/", controller.GetTodos)
    todo.Get("/:todoId", controller.GetTodo)
    todo.Put("/:todoId", controller.UpdateTodo)
    todo.Delete("/:todoId", controller.DeleteTodo)
}