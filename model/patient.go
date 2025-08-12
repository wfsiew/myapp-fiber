package model

import "encoding/xml"

type PatientDataResponse struct {
	Return string `xml:"return"`
}

type Return struct {
	XMLName xml.Name `xml:"Result"`
	Patient Patient  `xml:"Patient"`
}

type ReturnErr struct {
	XMLName xml.Name `xml:"Result"`
	Error   Error    `xml:"Error"`
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
	Prn            string         `xml:"PRN" json:"prn"`
	Name           Name           `xml:"Name" json:"name"`
	Resident       string         `xml:"Resident" json:"resident"`
	DOB            string         `xml:"DOB" json:"dob"`
	Sex            Sex            `xml:"Sex" json:"sex"`
	HomeAddress    HomeAddress    `xml:"HomeAddress" json:"homeAddress"`
	ContactNumber  ContactNumber  `xml:"ContactNumber" json:"contactNumber"`
	Document       []Document     `xml:"Document" json:"documents"`
	Nationality    Nationality    `xml:"Nationality" json:"nationality"`
	ChargeCategory ChargeCategory `xml:"ChargeCategory" json:"chargeCategory"`
	PaymentClass   PaymentClass   `xml:"PaymentClass" json:"paymentClass"`
}

type Error struct {
	ErrorCode    string `xml:"ErrorCode" json:"errorCode"`
	ErrorMessage string `xml:"ErrorMessage" json:"errorMessage"`
}
