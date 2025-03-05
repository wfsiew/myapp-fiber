package main

import (
	"app/database"
	_ "app/docs"
	"app/router"
	"app/utils"
	"fmt"
	"strings"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/django/v3"
	"github.com/swaggo/fiber-swagger"
	"github.com/valyala/fasttemplate"
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
    engine := django.New("./views", ".django")
    utils.SetValidator()
    utils.SetLogger()
    app := fiber.New(fiber.Config{
        Prefork: false,
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

    app.Get("/app/index", func(c *fiber.Ctx) error {
        // Render index
        err := c.Render("invoice.xml", fiber.Map{
            "inv": "INV1234598",
            "issue_date": "2024-05-28",
            "tin": "C5890633090",
            "brn": "200201024235",
        })
        c.Set(fiber.HeaderContentType, fiber.MIMETextXML)
        return err
    })
    app.Get("/app/test", func(c *fiber.Ctx) error {
        agent := fiber.Get("http://localhost:8000/app/index")
        _, body, _ := agent.Bytes()
        err := c.SendString(string(body[:]))
        c.Set(fiber.HeaderContentType, fiber.MIMETextXML)
        return err
    })
    app.Get("/app/inv", func(c *fiber.Ctx) error {
        s := GetInvoiceXml()
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

        t := fasttemplate.New(s, "{{", "}}")
        r := t.ExecuteString(map[string]interface{}{
            "inv": "INV1234598",
            "issue_date": "2024-05-28",
            "tin": "C5890633090",
            "brn": "200201024235",
            "additionalDocRef": GetAdditionalDocRef(ld),
            "data": ld,
        })
        c.Set(fiber.HeaderContentType, fiber.MIMETextXML)
        return c.SendString(r)
    })
    router.SetupRoutes(app)

    err := app.Listen(":8000")

    if err != nil {
        utils.Logger.Fatal().Err(err).Msg("Fiber app error")
    }
}

func GetAdditionalDocRef(ld []AdditionalDocumentReference) string {
    ls := make([]string, 0)
    for i := range ld {
        ls = append(ls, "<cac:AdditionalDocumentReference>\n")
        id := fmt.Sprintf("    <cbc:ID>%s</cbc:ID>\n", ld[i].ID)
        ls = append(ls, id)
        if ld[i].DocumentType != "" {
            dt := fmt.Sprintf("    <cbc:DocumentType>%s</cbc:DocumentType>\n", ld[i].DocumentType)
            ls = append(ls, dt)
        }

        if ld[i].DocumentDescription != "" {
            dd := fmt.Sprintf("    <cbc:DocumentDescription>%s</cbc:DocumentDescription>\n", ld[i].DocumentDescription)
            ls = append(ls, dd)
        }

        ls = append(ls, "</cac:AdditionalDocumentReference>\n")
    }

    s := strings.Join(ls, "\n")
    fmt.Println(s)
    return s
}

func GetInvoiceXml() string {
    s := `
<Invoice xmlns="urn:oasis:names:specification:ubl:schema:xsd:Invoice-2"
    xmlns:cac="urn:oasis:names:specification:ubl:schema:xsd:CommonAggregateComponents-2"
    xmlns:cbc="urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2">
    <cbc:ID>{{inv}}</cbc:ID>
    <cbc:IssueDate>{{issue_date}}</cbc:IssueDate>
    <cbc:IssueTime>15:30:00Z</cbc:IssueTime>
    <cbc:InvoiceTypeCode listVersionID="1.0">01</cbc:InvoiceTypeCode>
    <cbc:DocumentCurrencyCode>MYR</cbc:DocumentCurrencyCode>
    <cac:InvoicePeriod>
        <cbc:StartDate>2018-11-26</cbc:StartDate>
        <cbc:EndDate>2018-11-30</cbc:EndDate>
        <cbc:Description>Monthly</cbc:Description>
    </cac:InvoicePeriod>
    <cac:BillingReference>
        <cac:AdditionalDocumentReference>
            <cbc:ID>E12345678912</cbc:ID>
        </cac:AdditionalDocumentReference>
    </cac:BillingReference>
    {{additionalDocRef}}
    <cac:AccountingSupplierParty>
        <cbc:AdditionalAccountID schemeAgencyName="CertEX">CPT-CCN-W-211111-KL-000002</cbc:AdditionalAccountID>
        <cac:Party>
            <cbc:IndustryClassificationCode name="Growing of maize">01111</cbc:IndustryClassificationCode>
            <cac:PartyIdentification>
                <cbc:ID schemeID="TIN">{{ tin }}</cbc:ID>
            </cac:PartyIdentification>
            <cac:PartyIdentification>
                <cbc:ID schemeID="BRN">{{ brn }}</cbc:ID>
            </cac:PartyIdentification>
            <cac:PostalAddress>
                <cbc:CityName>Kuala Lumpur</cbc:CityName>
                <cbc:PostalZone>50480</cbc:PostalZone>
                <cbc:CountrySubentityCode>14</cbc:CountrySubentityCode>
                <cac:AddressLine>
                    <cbc:Line>Lot 66</cbc:Line>
                </cac:AddressLine>
                <cac:AddressLine>
                    <cbc:Line>Bangunan Merdeka</cbc:Line>
                </cac:AddressLine>
                <cac:AddressLine>
                    <cbc:Line>Persiaran Jaya</cbc:Line>
                </cac:AddressLine>
                <cac:Country>
                    <cbc:IdentificationCode listID="ISO3166-1" listAgencyID="6">MYS</cbc:IdentificationCode>
                </cac:Country>
            </cac:PostalAddress>
            <cac:PartyLegalEntity>
                <cbc:RegistrationName>AMS Setia Jaya Sdn. Bhd.</cbc:RegistrationName>
            </cac:PartyLegalEntity>
            <cac:Contact>
                <cbc:Telephone>+60-123456789</cbc:Telephone>
                <cbc:ElectronicMail>general.ams@supplier.com</cbc:ElectronicMail>
            </cac:Contact>
        </cac:Party>
    </cac:AccountingSupplierParty>
    <cac:AccountingCustomerParty>
        <cac:Party>
            <cac:PartyIdentification>
                <cbc:ID schemeID="TIN">C2584563200</cbc:ID>
            </cac:PartyIdentification>
            <cac:PartyIdentification>
                <cbc:ID schemeID="BRN">201901234567</cbc:ID>
            </cac:PartyIdentification>
            <cac:PostalAddress>
                <cbc:CityName>Kuala Lumpur</cbc:CityName>
                <cbc:PostalZone>50480</cbc:PostalZone>
                <cbc:CountrySubentityCode>14</cbc:CountrySubentityCode>
                <cac:AddressLine>
                    <cbc:Line>Lot 66</cbc:Line>
                </cac:AddressLine>
                <cac:AddressLine>
                    <cbc:Line>Bangunan Merdeka</cbc:Line>
                </cac:AddressLine>
                <cac:AddressLine>
                    <cbc:Line>Persiaran Jaya</cbc:Line>
                </cac:AddressLine>
                <cac:Country>
                    <cbc:IdentificationCode listID="ISO3166-1" listAgencyID="6">MYS</cbc:IdentificationCode>
                </cac:Country>
            </cac:PostalAddress>
            <cac:PartyLegalEntity>
                <cbc:RegistrationName>Hebat Group</cbc:RegistrationName>
            </cac:PartyLegalEntity>
            <cac:Contact>
                <cbc:Telephone>+60-123456789</cbc:Telephone>
                <cbc:ElectronicMail>name@buyer.com</cbc:ElectronicMail>
            </cac:Contact>
        </cac:Party>
    </cac:AccountingCustomerParty>
    <cac:Delivery>
        <cac:DeliveryParty>
            <cac:PartyIdentification>
                <cbc:ID schemeID="TIN">C2584563200</cbc:ID>
            </cac:PartyIdentification>
            <cac:PartyIdentification>
                <cbc:ID schemeID="BRN">201901234567</cbc:ID>
            </cac:PartyIdentification>
            <cac:PostalAddress>
                <cbc:CityName>Kuala Lumpur</cbc:CityName>
                <cbc:PostalZone>50480</cbc:PostalZone>
                <cbc:CountrySubentityCode>14</cbc:CountrySubentityCode>
                <cac:AddressLine>
                    <cbc:Line>Lot 66</cbc:Line>
                </cac:AddressLine>
                <cac:AddressLine>
                    <cbc:Line>Bangunan Merdeka</cbc:Line>
                </cac:AddressLine>
                <cac:AddressLine>
                    <cbc:Line>Persiaran Jaya</cbc:Line>
                </cac:AddressLine>
                <cac:Country>
                    <cbc:IdentificationCode listID="ISO3166-1" listAgencyID="6">MYS</cbc:IdentificationCode>
                </cac:Country>
            </cac:PostalAddress>
            <cac:PartyLegalEntity>
                <cbc:RegistrationName>Greenz Sdn. Bhd.</cbc:RegistrationName>
            </cac:PartyLegalEntity>
        </cac:DeliveryParty>
        <cac:Shipment>
            <cbc:ID>1234</cbc:ID>
            <cac:FreightAllowanceCharge>
                <cbc:ChargeIndicator>true</cbc:ChargeIndicator>
                <cbc:AllowanceChargeReason>Service charge</cbc:AllowanceChargeReason>
                <cbc:Amount currencyID="MYR">100</cbc:Amount>
            </cac:FreightAllowanceCharge>
        </cac:Shipment>
    </cac:Delivery>
    <cac:PaymentMeans>
        <cbc:PaymentMeansCode>01</cbc:PaymentMeansCode>
        <cac:PayeeFinancialAccount>
            <cbc:ID>1234567890123</cbc:ID>
        </cac:PayeeFinancialAccount>
    </cac:PaymentMeans>
    <cac:PaymentTerms>
        <cbc:Note>Payment method is cash</cbc:Note>
    </cac:PaymentTerms>
    <cac:PrepaidPayment>
        <cbc:ID>E12345678912</cbc:ID>
        <cbc:PaidAmount currencyID="MYR">1.00</cbc:PaidAmount>
        <cbc:PaidDate>2000-01-01</cbc:PaidDate>
        <cbc:PaidTime>12:00:00Z</cbc:PaidTime>
    </cac:PrepaidPayment>
    <cac:AllowanceCharge>
        <cbc:ChargeIndicator>false</cbc:ChargeIndicator>
        <cbc:AllowanceChargeReason>Sample Description</cbc:AllowanceChargeReason>
        <cbc:Amount currencyID="MYR">100</cbc:Amount>
    </cac:AllowanceCharge>
    <cac:AllowanceCharge>
        <cbc:ChargeIndicator>true</cbc:ChargeIndicator>
        <cbc:AllowanceChargeReason>Service charge</cbc:AllowanceChargeReason>
        <cbc:Amount currencyID="MYR">100</cbc:Amount>
    </cac:AllowanceCharge>
    <cac:TaxTotal>
        <cbc:TaxAmount currencyID="MYR">87.63</cbc:TaxAmount>
        <cac:TaxSubtotal>
            <cbc:TaxableAmount currencyID="MYR">87.63</cbc:TaxableAmount>
            <cbc:TaxAmount currencyID="MYR">87.63</cbc:TaxAmount>
            <cbc:Percent>10</cbc:Percent>
            <cbc:BaseUnitMeasure unitCode="C62">1</cbc:BaseUnitMeasure>
            <cbc:PerUnitAmount currencyID="MYR">10</cbc:PerUnitAmount>
            <cac:TaxCategory>
                <cbc:ID>01</cbc:ID>
                <cac:TaxScheme>
                    <cbc:ID schemeID="UN/ECE 5153" schemeAgencyID="6">OTH</cbc:ID>
                </cac:TaxScheme>
            </cac:TaxCategory>
        </cac:TaxSubtotal>
    </cac:TaxTotal>
    <cac:LegalMonetaryTotal>
        <cbc:LineExtensionAmount currencyID="MYR">1436.50</cbc:LineExtensionAmount>
        <cbc:TaxExclusiveAmount currencyID="MYR">1436.50</cbc:TaxExclusiveAmount>
        <cbc:TaxInclusiveAmount currencyID="MYR">1436.50</cbc:TaxInclusiveAmount>
        <cbc:AllowanceTotalAmount currencyID="MYR">1436.50</cbc:AllowanceTotalAmount>
        <cbc:ChargeTotalAmount currencyID="MYR">1436.50</cbc:ChargeTotalAmount>
        <cbc:PayableRoundingAmount currencyID="MYR">0.30</cbc:PayableRoundingAmount>
        <cbc:PayableAmount currencyID="MYR">1436.50</cbc:PayableAmount>
    </cac:LegalMonetaryTotal>
    <cac:InvoiceLine>
        <cbc:ID>1</cbc:ID>
        <cbc:InvoicedQuantity unitCode="C62">1</cbc:InvoicedQuantity>
        <cbc:LineExtensionAmount currencyID="MYR">1436.50</cbc:LineExtensionAmount>
        <cac:AllowanceCharge>
            <cbc:ChargeIndicator>false</cbc:ChargeIndicator>
            <cbc:AllowanceChargeReason>Sample Description</cbc:AllowanceChargeReason>
            <cbc:MultiplierFactorNumeric>0.15</cbc:MultiplierFactorNumeric>
            <cbc:Amount currencyID="MYR">100</cbc:Amount>
        </cac:AllowanceCharge>
        <cac:AllowanceCharge>
            <cbc:ChargeIndicator>true</cbc:ChargeIndicator>
            <cbc:AllowanceChargeReason>Sample Description</cbc:AllowanceChargeReason>
            <cbc:MultiplierFactorNumeric>0.10</cbc:MultiplierFactorNumeric>
            <cbc:Amount currencyID="MYR">100</cbc:Amount>
        </cac:AllowanceCharge>
        <cac:TaxTotal>
            <cbc:TaxAmount currencyID="MYR">0</cbc:TaxAmount>
            <cac:TaxSubtotal>
                <cbc:TaxableAmount currencyID="MYR">1460.50</cbc:TaxableAmount>
                <cbc:TaxAmount currencyID="MYR">0</cbc:TaxAmount>
                <cac:TaxCategory>
                    <cbc:ID>E</cbc:ID>
                    <cbc:Percent>6.00</cbc:Percent>
                    <cbc:TaxExemptionReason>Exempt New Means of Transport</cbc:TaxExemptionReason>
                    <cac:TaxScheme>
                        <cbc:ID schemeID="UN/ECE 5153" schemeAgencyID="6">OTH</cbc:ID>
                    </cac:TaxScheme>
                </cac:TaxCategory>
            </cac:TaxSubtotal>
        </cac:TaxTotal>
        <cac:Item>
            <cbc:Description>Laptop Peripherals</cbc:Description>
            <cac:OriginCountry>
                <cbc:IdentificationCode>MYS</cbc:IdentificationCode>
            </cac:OriginCountry>
            <cac:CommodityClassification>
                <cbc:ItemClassificationCode listID="PTC">001</cbc:ItemClassificationCode>
            </cac:CommodityClassification>
            <cac:CommodityClassification>
                <cbc:ItemClassificationCode listID="CLASS">002</cbc:ItemClassificationCode>
            </cac:CommodityClassification>
        </cac:Item>
        <cac:Price>
            <cbc:PriceAmount currencyID="MYR">17</cbc:PriceAmount>
        </cac:Price>
        <cac:ItemPriceExtension>
            <cbc:Amount currencyID="MYR">100</cbc:Amount>
        </cac:ItemPriceExtension>
    </cac:InvoiceLine>
</Invoice>
    `
    return s
}