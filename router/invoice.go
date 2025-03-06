package router

import (
    "app/controller"
    "github.com/gofiber/fiber/v2"
)

func SetupInvoiceRoutes(router fiber.Router) {
    inv := router.Group("/invoice")
    inv.Get("/index", controller.GetIndex)
    inv.Get("/data", controller.GetInvData)
}