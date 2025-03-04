package main

import (
	"app/database"
	_ "app/docs"
	"app/router"
	"app/utils"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/swaggo/fiber-swagger"
	"github.com/gofiber/template/django/v3"
)

// @title Swagger App API
// @version 1.0
// @description This is a sample server Petstore server.
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @BasePath /app
func main() {
	engine := django.New("./views", ".django")
	utils.SetValidator()
	utils.SetLogger()
	app := fiber.New(fiber.Config{
		Prefork: true,
		Views: engine,
	})
	app.Use(recover.New())
	app.Use(compress.New())
	// app.Use("/app", cors.Config{AllowOrigins: "*"})
	app.Use(fiberzerolog.New(fiberzerolog.Config{
        Logger: &utils.Logger,
    }))
	database.ConnectDB()
	defer database.GetDb().Close()

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index
		return c.Render("file", fiber.Map{
			"Title": "Hello, World!",
		})
	})
	router.SetupRoutes(app)

	err := app.Listen(":8000")

	if err != nil {
		utils.Logger.Fatal().Err(err).Msg("Fiber app error")
	}
}