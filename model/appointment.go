package model

import "encoding/xml"

type FutureAppointmentResponse struct {
	Return string `xml:"return"`
}

type

type Appointment struct {
    AppointmentNumber string `xml:"AppointmentNumber" json:"appointmentNumber"`
    Date              string `xml:"Date" json:"date"`
    Day               string `xml:"Day" json:"day"`
    StartTime         string `xml:"StartTime" json:"startTime"`
    EndTime           string `xml:"EndTime" json:"endTime"`
    DoctorMCR         string `xml:"DoctorMCR" json:"doctorMcr"`
    DoctorName        string `xml:"DoctorName" json:"doctorName"`
    SpecialtyCode     string `xml:"SpecialtyCode" json:"specialtyCode"`
    Specialty         string `xml:"Specialty" json:"specialty"`
    Clinic            string `xml:"Clinic" json:"clinic"`
    Room              string `xml:"Room" json:"room"`
    CaseType          string `xml:"CaseType" json:"caseType"`
}
