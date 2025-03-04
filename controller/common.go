package controller

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"app/database"
	"app/model"
	"app/utils"
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
