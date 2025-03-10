package controller

import (
    "fmt"
    "encoding/xml"
    "app/config"
    "app/model"
    "github.com/gofiber/fiber/v2"
)

// GetIndex
//
// @Tags Invoice
// @Produce xml
// @Success 200 
// @Router /app/invoice/index [get]
func GetIndex(c *fiber.Ctx) error {
    ld := make([]model.AdditionalDocumentReference, 0)
    ld = append(ld, model.AdditionalDocumentReference{
        ID: "E12345678912",
        DocumentType: "CustomsImportForm",
    })
    ld = append(ld, model.AdditionalDocumentReference{
        ID: "ASEAN-Australia-New Zealand FTA (AANZFTA)",
        DocumentType: "FreeTradeAgreement",
        DocumentDescription: "Sample Description",
    })
    ld = append(ld, model.AdditionalDocumentReference{
        ID: "E12345678912",
        DocumentType: "K2",
    })
    ld = append(ld, model.AdditionalDocumentReference{
        ID: "CIF",
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
}

// GetInvData
//
// @Tags Invoice
// @Produce xml
// @Success 200 
// @Router /app/invoice/data [get]
func GetInvData(c *fiber.Ctx) error {
    port := config.Config("PORT")
    agent := fiber.Get(fmt.Sprintf("http://localhost:%s/app/invoice/index", port))
    _, body, _ := agent.Bytes()
    err := c.SendString(string(body[:]))
    c.Set(fiber.HeaderContentType, fiber.MIMETextXML)
    return err
}

func GetInvTest(c *fiber.Ctx) error {
    x := &model.InvoicePeriod{}
    x.StartDate = "2018-11-26"
    x.EndDate = "2018-11-30"
    x.Description = "Monthly"
    o, _ := xml.MarshalIndent(x, " ", "  ")
    fmt.Println(string(o))
    start := `<Invoice xmlns="urn:oasis:names:specification:ubl:schema:xsd:Invoice-2"
    xmlns:cac="urn:oasis:names:specification:ubl:schema:xsd:CommonAggregateComponents-2"
    xmlns:cbc="urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2">
    `
    end := "</Invoice>"
    err := c.SendString(start + string(o) + end)
    c.Set(fiber.HeaderContentType, fiber.MIMETextXML)
    return err
}