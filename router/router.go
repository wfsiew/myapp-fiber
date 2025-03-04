package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/app", logger.New())
	SetupAuthRoutes(api)
	SetupTodoRoutes(api)
	SetupCommonRoutes(api)
}