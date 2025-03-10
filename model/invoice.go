package model

type InvoicePeriod struct {
    StartDate string `json:"startDate"`
    EndDate string `json:"endDate"`
    Description string `json:"description"`
}

type BillingReference struct {
    AdditionalDocumentReference AdditionalDocumentReference `json:"additionalDocumentReference"`
}

type AdditionalDocumentReference struct {
    ID string `json:"id"`
    DocumentType string `json:"documentType"`
    DocumentDescription string `json:"documentDescription"`
}

type Contact struct {
    Telephone string `json:"telephone"`
    ElectronicMail string `json:"electronicMail"`
}

type PartyLegalEntity struct {
    RegistrationName string `json:"registrationName"`
}

type AddressLine struct {
    Line string `json:"line"`
}

type Country struct {
    IdentificationCode string `json:"identificationCode"`
    AttrlistID string `json:"listId"`
    AttrlistAgencyID string `json:"listAgencyId"`
}

type PostalAddress struct {
    CityName string `json:"cityName"`
    PostalZone string `json:"postalZone"`
    CountrySubentityCode string `json:"countrySubentityCode"`
    AddressLineList []AddressLine `json:"addressLineList"`
    Country Country `json:"country"`
}

type PartyIdentification struct {
    ID string `json:"id"`
    AttrschemeID string `json:"schemeId"`
}

type Party struct {
    IndustryClassificationCode string `json:"industryClassificationCode"`
    Attrname string `json:"name"`
    PartyIdentificationList []PartyIdentification `json:"partyIdentificationList"`
    PostalAddress PostalAddress `json:"postalAddress"`
    PartyLegalEntity PartyLegalEntity `json:"partyLegalEntity"`
    Contact Contact `json:"contact"`
}

type AccountingSupplierParty struct {
    AdditionalAccountID string `json:"additionalAccountID"`
    AttrschemeAgencyName string `json:"schemeAgencyName"`
    Party Party `json:"party"`
}

type AccountingCustomerParty struct {
    Party Party `json:"party"`
}

type DeliveryParty struct {
    PartyIdentificationList []PartyIdentification `json:"partyIdentificationList"`
    PostalAddress PostalAddress `json:"postalAddress"`
    PartyLegalEntity PartyLegalEntity `json:"partyLegalEntity"`
}

type FreightAllowanceCharge struct {
    ChargeIndicator bool `json:"chargeIndicator"`
    AllowanceChargeReason string `json:"allowanceChargeReason"`
    Amount string `json:"amount"`
    AttrcurrencyID string `json:"currencyId"`
}

type Shipment struct {
    ID string `json:"id"`
    FreightAllowanceCharge FreightAllowanceCharge `json:"freightAllowanceCharge"`
}

type Delivery struct {
    DeliveryParty DeliveryParty `json:"deliveryParty"`
    Shipment Shipment `json:"shipment"`
}

type PayeeFinancialAccount struct {
    ID string `json:"id"`
}

type PaymentMeans struct {
    PaymentMeansCode string `json:"paymentMeansCode"`
    PayeeFinancialAccount PayeeFinancialAccount `json:"payeeFinancialAccount"`
}

type PaymentTerms struct {
    Note string `json:"note"`
}

type PrepaidPayment struct {
    ID string `json:"id"`
    PaidAmount string `json:"paidAmount"`
    AttrcurrencyID string `json:"attrcurrencyId"`
    PaidDate string `json:"paidDate"`
    PaidTime string `json:"paidTime"`
}

type AllowanceCharge struct {
    ChargeIndicator bool `json:"chargeIndicator"`
    AllowanceChargeReason string `json:"allowanceChargeReason"`
    MultiplierFactorNumeric string `json:"multiplierFactorNumeric"`
    Amount string `json:"amount"`
    AttrcurrencyID string `json:"currencyId"`
}

type TaxScheme struct {
    ID string `json:"id"`
    AttrschemeID string `json:"schemeId"`
    AttrschemeAgencyID string `json:"schemeAgencyId"`
}

type TaxCategory struct {
    ID string `json:"id"`
    Percent string `json:"percent"`
    TaxExemptionReason string `json:"taxExemptionReason"`
    TaxScheme TaxScheme `json:"taxScheme"`
}

type TaxSubtotal struct {
    TaxableAmount string `json:"taxableAmount"`
    AttrcurrencyID string `json:"currencyId"`
    TaxAmount string `json:"taxAmount"`
    AttrcurrencyID1 string `json:"currencyId1"`
    Percent string `json:"percent"`
    BaseUnitMeasure string `json:"baseUnitMeasure"`
    AttrunitCode string `json:"unitCode"`
    PerUnitAmount string `json:"perUnitAmount"`
    AttrcurrencyID2 string `json:"currencyId2"`
    TaxCategory TaxCategory `json:"taxCategory"`
}

type TaxTotal struct {
    TaxAmount string `json:"taxAmount"`
    AttrcurrencyID string `json:"currencyId"`
    TaxSubtotal TaxSubtotal `json:"taxSubtotal"`
}

type LegalMonetaryTotal struct {
    LineExtensionAmount string `json:"lineExtensionAmount"`
    AttrcurrencyID string `json:"currencyId"`
    TaxExclusiveAmount string `json:"taxExclusiveAmount"`
    AttrcurrencyID1 string `json:"currencyId1"`
    TaxInclusiveAmount string `json:"taxInclusiveAmount"`
    AttrcurrencyID2 string `json:"currencyId2"`
    AllowanceTotalAmount string `json:"allowanceTotalAmount"`
    AttrcurrencyID3 string `json:"currencyId3"`
    ChargeTotalAmount string `json:"chargeTotalAmount"`
    AttrcurrencyID4 string `json:"currencyId4"`
    PayableRoundingAmount string `json:"payableRoundingAmount"`
    AttrcurrencyID5 string `json:"currencyId5"`
    PayableAmount string `json:"payableAmount"`
}

type CommodityClassification struct {
    ItemClassificationCode string `json:"itemClassificationCode"`
    AttrlistID string `json:"listId"`
}

type OriginCountry struct {
    IdentificationCode string `json:"identificationCode"`
}

type Item struct {
    Description string `json:"description"`
    OriginCountry OriginCountry `json:"originCountry"`
    CommodityClassificationList []CommodityClassification `json:"commodityClassificationList"`
}

type Price struct {
    PriceAmount string `json:"amount"`
    AttrcurrencyID string `json:"currencyId"`
}

type ItemPriceExtension struct {
    Amount string `json:"amount"`
    AttrcurrencyID string `json:"currencyId"`
}

type InvoiceLine struct {
    ID string `json:"id"`
    InvoicedQuantity string `json:"invoicedQuantity"`
    AttrunitCode string `json:"unitCode"`
    LineExtensionAmount string `json:"lineExtensionAmount"`
    AttrcurrencyID string `json:"currencyId"`
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