package controller

import (
	"app/database"
	"app/model"
	"app/utils"
	"crypto/tls"
	"database/sql"
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"siteminds.dev/gosoap"
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
		ccode   sql.NullString
		cname   sql.NullString
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
			TelCode:     telcode.String,
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
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	url := "https://nh-europa.nova-vesalius.com/ihp_uat/web_services/AUTHENTICATION/Login.cfc?WSDL"
	httpClient := &http.Client{
		Timeout:   1500 * time.Millisecond,
		Transport: tr,
	}

	config := gosoap.Config{
		Endpoint: "https://nh-europa.nova-vesalius.com/ihp_uat/web_services/AUTHENTICATION/Login.cfc",
	}
	soap, err := gosoap.SoapClientWithConfig(url, httpClient, &config)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "ERROR",
			"message": fmt.Errorf("SoapClient: %v", err).Error(),
		})
	}

	params := gosoap.Params{
		"company_code": "IHP",
		"system_code":  "MOBILE",
		"password":     "password",
	}

	res, err := soap.Call("login", params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "ERROR",
			"message": fmt.Errorf("SoapCall: %v", err).Error(),
		})
	}

	var r model.LoginResponse
	res.Unmarshal(&r)
	s := strings.ReplaceAll(r.Return, `encoding="UTF-8">`, `encoding="UTF-8"?>`)
	fmt.Println(s)
	result := model.LoginReturn{}
	resulterr := model.ReturnErr{}
	err = xml.Unmarshal([]byte(s), &result)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "ERROR",
			"message": fmt.Errorf("unmarshal return result: %v", err).Error(),
		})
	}

	if result.Token.Token_number == "" {
		err = xml.Unmarshal([]byte(s), &resulterr)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"code":    "ERROR",
				"message": fmt.Errorf("unmarshal return resultErr: %v", err).Error(),
			})
		}

		return c.JSON(resulterr.Error)
	}

	return c.JSON(result.Token)
}

// Logout
//
// @Tags Common
// @Produce json
// @Param        token   path      string  true  "token"
// @Success 200 {object} model.Success
// @Router /app/common/logout/{token} [get]
func Logout(c *fiber.Ctx) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	url := "https://nh-europa.nova-vesalius.com/ihp_uat/web_services/AUTHENTICATION/Logout.cfc?WSDL"
	httpClient := &http.Client{
		Timeout:   1500 * time.Millisecond,
		Transport: tr,
	}

	config := gosoap.Config{
		Endpoint: "https://nh-europa.nova-vesalius.com/ihp_uat/web_services/AUTHENTICATION/Logout.cfc",
	}
	soap, err := gosoap.SoapClientWithConfig(url, httpClient, &config)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "ERROR",
			"message": fmt.Errorf("SoapClient: %v", err).Error(),
		})
	}

	params := gosoap.Params{
		"token_number": c.Params("token"),
	}

	res, err := soap.Call("Logout", params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "ERROR",
			"message": fmt.Errorf("SoapCall: %v", err).Error(),
		})
	}

	var r model.LoginResponse
	res.Unmarshal(&r)
	result := model.LogoutReturn{}
	resulterr := model.ReturnErr{}
	err = xml.Unmarshal([]byte(r.Return), &result)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "ERROR",
			"message": fmt.Errorf("unmarshal return: %v", err).Error(),
		})
	}

	if result.Success.Code == "" {
		err = xml.Unmarshal([]byte(r.Return), &resulterr)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"code":    "ERROR",
				"message": fmt.Errorf("unmarshal return: %v", err).Error(),
			})
		}

		return c.JSON(resulterr.Error)
	}

	return c.JSON(result.Success)
}

// GetPatientData
//
// @Tags Common
// @Produce json
// @Success 200 {object} model.Patient
// @Router /app/common/patient [get]
func GetPatientData(c *fiber.Ctx) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	url := "https://nh-europa.nova-vesalius.com/ihp_uat/web_services/PATIENT/GetPatientData.cfc?WSDL"
	httpClient := &http.Client{
		Timeout:   1500 * time.Millisecond,
		Transport: tr,
	}

	config := gosoap.Config{
		Endpoint: "https://nh-europa.nova-vesalius.com/ihp_uat/web_services/PATIENT/GetPatientData.cfc",
	}
	soap, err := gosoap.SoapClientWithConfig(url, httpClient, &config)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "ERROR",
			"message": fmt.Errorf("SoapClient: %v", err).Error(),
		})
	}

	params := gosoap.Params{
		"prn": "20-000110",
	}

	res, err := soap.Call("getPatientData", params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "ERROR",
			"message": fmt.Errorf("SoapCall: %v", err).Error(),
		})
	}

	var r model.PatientDataResponse
	res.Unmarshal(&r)
	result := model.Return{}
	resulterr := model.ReturnErr{}
	err = xml.Unmarshal([]byte(r.Return), &result)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "ERROR",
			"message": fmt.Errorf("unmarshal return: %v", err).Error(),
		})
	}

	if result.Patient.Prn == "" {
		err = xml.Unmarshal([]byte(r.Return), &resulterr)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"code":    "ERROR",
				"message": fmt.Errorf("unmarshal return: %v", err).Error(),
			})
		}

		return c.JSON(resulterr.Error)
	}

	return c.JSON(result.Patient)
}
