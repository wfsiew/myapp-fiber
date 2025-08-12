package controller

import (
	"app/database"
	"app/model"
	"app/utils"
	"database/sql"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/tiaguinho/gosoap"
)

// GetCountries
//
// @Tags Common
// @Produce json
// @Success 200 {array} model.CountryData
// @Router /app/common/country/list [get]
func GetCountries(c *fiber.Ctx) error {
    db := database.GetDb()
    if db == nil {
        utils.ILogger.Info().Msg("db is nil")
        return c.JSON([]model.CountryData{})
    }

    rows, err := db.Query("SELECT COUNTRY_CODE, COUNTRY_NAME, TEL_CODE FROM NOVA_COUNTRY")
    if err != nil {
        if rows != nil {
            defer rows.Close()
        }
        
        utils.Logger.Err(err).Msg(err.Error())
        return c.JSON([]model.CountryData{})
    }

    defer rows.Close()

    var (
        ccode sql.NullString
        cname sql.NullString
        telcode sql.NullString
    )
    var ls []model.CountryData = make([]model.CountryData, 0)

    for rows.Next() {
        err := rows.Scan(&ccode, &cname, &telcode)

        if err != nil {
            utils.Logger.Err(err).Msg(err.Error())
            continue
        }

        o := model.CountryData{
            CountryCode: ccode.String,
            CountryName: cname.String,
            TelCode: telcode.String,
        }
        ls = append(ls, o)
    }

    return c.JSON(ls)
}

// Authenticate
//
// @Tags Common
// @Produce json
// @Success 200 {object} model.Token
// @Router /app/common/login [get]
func Authenticate(c *fiber.Ctx) error {
    url := "https://nh-europa.nova-vesalius.com/ihp_uat/web_services/AUTHENTICATION/Login.cfc?WSDL"
    httpClient := &http.Client{
		Timeout: 1500 * time.Millisecond,
	}

    soap, err := gosoap.SoapClient(url, httpClient)
	if err != nil {
		log.Fatalf("SoapClient error: %s", err)
	}

    params := gosoap.Params{
		"company_code": "IHP",
        "system_code": "MOBILE",
        "password": "password",
	}

    res, err := soap.Call("login", params)
	if err != nil {
		log.Fatalf("Call error: %s", err)
	}

    var r model.LoginResponse
    res.Unmarshal(&r)
    result := model.LoginReturn{}
    resulterr := model.ReturnErr{}
    fmt.Println(r.Return)
    err = xml.Unmarshal([]byte(r.Return), &result)
    if err != nil {
		log.Fatalf("xml.Unmarshal error: %s", err)
	}

    if result.Token.Token_number == "" {
        err = xml.Unmarshal([]byte(r.Return), &resulterr)
        if err != nil {
            log.Fatalf("xml.Unmarshal error: %s", err)
        }

        return c.JSON(resulterr.Error)
    }

    return c.JSON(result.Token)
}

// GetPatientData
//
// @Tags Common
// @Produce json
// @Success 200 {object} model.Patient
// @Router /app/common/patient [get]
func GetPatientData(c *fiber.Ctx) error {
    url := "https://nh-europa.nova-vesalius.com/ihp_uat/web_services/PATIENT/GetPatientData.cfc?WSDL"
	httpClient := &http.Client{
		Timeout: 1500 * time.Millisecond,
	}

    soap, err := gosoap.SoapClient(url, httpClient)
	if err != nil {
		log.Fatalf("SoapClient error: %s", err)
	}

	params := gosoap.Params{
		"prn": "20-000110",
	}

    res, err := soap.Call("getPatientData", params)
	if err != nil {
		log.Fatalf("Call error: %s", err)
	}

    var r model.PatientDataResponse
	res.Unmarshal(&r)
	result := model.Return{}
    resulterr := model.ReturnErr{}
	err = xml.Unmarshal([]byte(r.Return), &result)
	if err != nil {
		log.Fatalf("xml.Unmarshal error: %s", err)
	}

    if result.Patient.Prn == "" {
        err = xml.Unmarshal([]byte(r.Return), &resulterr)
        if err != nil {
            log.Fatalf("xml.Unmarshal error: %s", err)
        }

        return c.JSON(resulterr.Error)
    }

    return c.JSON(result.Patient)
}
