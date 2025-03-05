package main

import (
	"app/config"
	"app/database"
	_ "app/docs"
	"app/router"
	"app/utils"
	"fmt"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/django/v3"
	"github.com/swaggo/fiber-swagger"
)

type AdditionalDocumentReference struct {
    ID string
    DocumentType string
    DocumentDescription string
}

// @title Swagger App API
// @version 1.0
// @description This is a sample server Petstore server.
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @BasePath /app
func main() {
    port := config.Config("PORT")
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

    app.Get("/app/index", logger.New(), func(c *fiber.Ctx) error {
        // Render index
        ld := make([]AdditionalDocumentReference, 0)
        ld = append(ld, AdditionalDocumentReference{
            ID: "E12345678912",
            DocumentType: "CustomsImportForm",
            DocumentDescription: "",
        })
        ld = append(ld, AdditionalDocumentReference{
            ID: "ASEAN-Australia-New Zealand FTA (AANZFTA)",
            DocumentType: "FreeTradeAgreement",
            DocumentDescription: "Sample Description",
        })
        ld = append(ld, AdditionalDocumentReference{
            ID: "E12345678912",
            DocumentType: "K2",
            DocumentDescription: "",
        })
        ld = append(ld, AdditionalDocumentReference{
            ID: "CIF",
            DocumentType: "",
            DocumentDescription: "",
        })
        err := c.Render("invoice.xml", fiber.Map{
            "inv": "INV1234598",
            "issue_date": "2024-05-28",
            "tin": "C5890633090",
            "brn": "200201024235",
            "zdditionalDocumentReferenceList": ld,
        })
        c.Set(fiber.HeaderContentType, fiber.MIMETextXML)
        return err
    })
    app.Get("/app/inv", logger.New(), func(c *fiber.Ctx) error {
        agent := fiber.Get(fmt.Sprintf("http://localhost:%s/app/index", port))
        _, body, _ := agent.Bytes()
        err := c.SendString(string(body[:]))
        c.Set(fiber.HeaderContentType, fiber.MIMETextXML)
        return err
    })
    router.SetupRoutes(app)

    err := app.Listen(fmt.Sprintf(":%s", port))

    if err != nil {
        utils.Logger.Fatal().Err(err).Msg("Fiber app error")
    }
}