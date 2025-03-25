package model

import "encoding/xml"

type PatientDataResponse struct {
	Return string `xml:"return"`
}

type Return struct {
	XMLName xml.Name `xml:"Result"`
	Patient Patient  `xml:"Patient"`
}

type Name struct {
	Title      string `xml:"Title"`
	FirstName  string `xml:"FirstName"`
	MiddleName string `xml:"MiddleName"`
	LastName   string `xml:"LastName"`
}

type Sex struct {
	Code        string `xml:"Code"`
	Description string `xml:"Description"`
}

type HomeAddress struct {
	Address1   string `xml:"Address1"`
	Address2   string `xml:"Address2"`
	Address3   string `xml:"Address3"`
	CityState  string `xml:"CityState"`
	PostalCode string `xml:"PostalCode"`
	Country    string `xml:"Country"`
}

type ContactNumber struct {
	Home  string `xml:"Home"`
	Email string `xml:"Email"`
}

type Document struct {
	Code        string `xml:"Code"`
	Description string `xml:"Description"`
	Value       string `xml:"Value"`
	ExpireDate  string `xml:"ExpireDate"`
}

type Nationality struct {
	Code        string `xml:"Code"`
	Description string `xml:"Description"`
}

type ChargeCategory struct {
	Code        string `xml:"Code"`
	Description string `xml:"Description"`
}

type PaymentClass struct {
	Code        string `xml:"Code"`
	Description string `xml:"Description"`
}

type Patient struct {
	Prn            string         `xml:"PRN"`
	Name           Name           `xml:"Name"`
	Resident       string         `xml:"Resident"`
	DOB            string         `xml:"DOB"`
	Sex            Sex            `xml:"Sex"`
	HomeAddress    HomeAddress    `xml:"HomeAddress"`
	ContactNumber  ContactNumber  `xml:"ContactNumber"`
	Document       []Document     `xml:"Document"`
	Nationality    Nationality    `xml:"Nationality"`
	ChargeCategory ChargeCategory `xml:"ChargeCategory"`
	PaymentClass   PaymentClass   `xml:"PaymentClass"`
}
