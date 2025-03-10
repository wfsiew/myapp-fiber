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
    AttrlistID string
    AttrlistAgencyID string
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
    Attrname string
    PartyIdentificationList []PartyIdentification
    PostalAddress PostalAddress
    PartyLegalEntity PartyLegalEntity
    Contact Contact
}

type AccountingSupplierParty struct {
    AdditionalAccountID string
    AttrschemeAgencyName string
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
    AttrcurrencyID string
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
    AttrcurrencyID string
    PaidDate string
    PaidTime string
}

type AllowanceCharge struct {
    ChargeIndicator bool
    AllowanceChargeReason string
    MultiplierFactorNumeric string
    Amount string
    AttrcurrencyID string
}

type TaxScheme struct {
    ID string
    AttrschemeID string
    AttrschemeAgencyID string
}

type TaxCategory struct {
    ID string
    Percent string
    TaxExemptionReason string
    TaxScheme TaxScheme
}

type TaxSubtotal struct {
    TaxableAmount string
    AttrcurrencyID string
    TaxAmount string
    AttrcurrencyID1 string
    Percent string
    BaseUnitMeasure string
    AttrunitCode string
    PerUnitAmount string
    AttrcurrencyID2 string
    TaxCategory TaxCategory
}

type TaxTotal struct {
    TaxAmount string
    AttrcurrencyID string
    TaxSubtotal TaxSubtotal
}

type LegalMonetaryTotal struct {
    LineExtensionAmount string
    AttrcurrencyID string
    TaxExclusiveAmount string
    AttrcurrencyID1 string
    TaxInclusiveAmount string
    AttrcurrencyID2 string
    AllowanceTotalAmount string
    AttrcurrencyID3 string
    ChargeTotalAmount string
    AttrcurrencyID4 string
    PayableRoundingAmount string
    AttrcurrencyID5 string
    PayableAmount string
}

type CommodityClassification struct {
    ItemClassificationCode string
    AttrlistID string
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
    AttrcurrencyID string
}

type ItemPriceExtension struct {
    Amount string
    AttrcurrencyID string
}

type InvoiceLine struct {
    ID string
    InvoicedQuantity string
    AttrunitCode string
    LineExtensionAmount string
    AttrcurrencyID string
}

type Invoice struct {
    ID string
    IssueDate string
    IssueTime string
    InvoiceTypeCode string
    AttrlistVersionID string
    DocumentCurrencyCode string
    InvoicePeriod InvoicePeriod
    BillingReference BillingReference
    AdditionalDocumentReference []AdditionalDocumentReference
    AccountingSupplierParty AccountingSupplierParty
    AccountingCustomerParty AccountingCustomerParty
    Delivery Delivery
    PaymentMeans PaymentMeans
    PaymentTerms PaymentTerms
    PrepaidPayment PrepaidPayment
    AllowanceCharge []AllowanceCharge
    TaxTotal TaxTotal
    LegalMonetaryTotal LegalMonetaryTotal
    InvoiceLine InvoiceLine
}