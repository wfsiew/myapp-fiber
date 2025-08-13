package model

import "encoding/xml"

type LoginResponse struct {
    Return string `xml:"return"`
}

type LoginReturn struct {
    XMLName xml.Name `xml:"Result"`
    Token   Token    `xml:"Token"`
}

type Token struct {
    Token_number string `xml:"Token_number" json:"tokenNumber"`
}

type LogoutReturn struct {
    XMLName xml.Name `xml:"Result"`
    Success Success  `xml:"Success"`
}

type Success struct {
    Code    string `xml:"Code" json:"code"`
    Message string `xml:"Message" json:"message"`
}
