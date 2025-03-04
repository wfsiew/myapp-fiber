package router

import (
    "app/controller"
    "github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(router fiber.Router) {
    auth := router.Group("/auth")
    auth.Post("/login", controller.Login)
    auth.Post("/register", controller.Register)
}