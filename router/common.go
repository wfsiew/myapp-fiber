package router

import (
    "app/controller"
    "github.com/gofiber/fiber/v2"
)

func SetupCommonRoutes(router fiber.Router) {
    common := router.Group("/common")
    common.Get("/country/list", controller.GetCountries)
    common.Get("/login", controller.Authenticate)
    common.Get("/patient", controller.GetPatientData)
}