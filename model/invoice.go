package model

type InvoicePeriod struct {
    StartDate string
    EndDate string
    Description string
}

type BillingReference struct {
    AdditionalDocumentReference AdditionalDocumentReference
}

type AdditionalDocumentReference struct {
    ID string
    DocumentType string
    DocumentDescription string
}

type Contact struct {
    Telephone string
    ElectronicMail string
}

type PartyLegalEntity struct {
    RegistrationName string
}

type AddressLine struct {
    Line string
}

type Country struct {
    IdentificationCode string
    MlistID string
    MlistAgencyID string
}

type PostalAddress struct {
    CityName string
    PostalZone string
    CountrySubentityCode string
    AddressLineList []AddressLine
    Country Country
}

type PartyIdentification struct {
    ID string
    MschemeID string
}

type Party struct {
    IndustryClassificationCode string
    Mname string
    PartyIdentificationList []PartyIdentification
    PostalAddress PostalAddress
    PartyLegalEntity PartyLegalEntity
    Contact Contact
}

type AccountingSupplierParty struct {
    AdditionalAccountID string
    MschemeAgencyName string
    Party Party
}

type AccountingCustomerParty struct {
    Party Party
}

type DeliveryParty struct {
    PartyIdentificationList []PartyIdentification
    PostalAddress PostalAddress
    PartyLegalEntity PartyLegalEntity
}

type FreightAllowanceCharge struct {
    ChargeIndicator bool
    AllowanceChargeReason string
    Amount string
    McurrencyID string
}

type Shipment struct {
    ID string
    FreightAllowanceCharge FreightAllowanceCharge
}

type Delivery struct {
    DeliveryParty DeliveryParty
    Shipment Shipment
}

type PayeeFinancialAccount struct {
    ID string
}

type PaymentMeans struct {
    PaymentMeansCode string
    PayeeFinancialAccount PayeeFinancialAccount
}

type PaymentTerms struct {
    Note string
}

type PrepaidPayment struct {
    ID string
    PaidAmount string
    McurrencyID string
    PaidDate string
    PaidTime string
}

type AllowanceCharge struct {
    ChargeIndicator bool
    AllowanceChargeReason string
    MultiplierFactorNumeric string
    Amount string
    McurrencyID string
}

type TaxScheme struct {
    ID string
    MschemeID string
    MschemeAgencyID string
}

type TaxCategory struct {
    ID string
    Percent string
    TaxExemptionReason string
    TaxScheme TaxScheme
}

type TaxSubtotal struct {
    TaxableAmount string
    McurrencyID string
    TaxAmount string
    McurrencyID1 string
    Percent string
    BaseUnitMeasure string
    MunitCode string
    PerUnitAmount string
    McurrencyID2 string
    TaxCategory TaxCategory
}

type TaxTotal struct {
    TaxAmount string
    McurrencyID string
    TaxSubtotal TaxSubtotal
}

type LegalMonetaryTotal struct {
    LineExtensionAmount string
    McurrencyID string
    TaxExclusiveAmount string
    McurrencyID1 string
    TaxInclusiveAmount string
    McurrencyID2 string
    AllowanceTotalAmount string
    McurrencyID3 string
    ChargeTotalAmount string
    McurrencyID4 string
    PayableRoundingAmount string
    McurrencyID5 string
    PayableAmount string
}

type CommodityClassification struct {
    ItemClassificationCode string
    MlistID string
}

type OriginCountry struct {
    IdentificationCode string
}

type Item struct {
    Description string
    OriginCountry OriginCountry
    CommodityClassification []CommodityClassification
}

type Price struct {
    PriceAmount string
    McurrencyID string
}

type ItemPriceExtension struct {
    Amount string
    McurrencyID string
}

type InvoiceLine struct {
    ID string
    InvoicedQuantity string
    MunitCode string
    LineExtensionAmount string
    McurrencyID string
}