package controllers

import (
	"database/sql"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/helpnumber/geo/georeverse"
	"github.com/kisielk/sqlstruct"
)

func GetPhonenumbers(c *gin.Context) {
	db := c.MustGet("database").(*sql.DB)
	language_code := strings.ToLower(c.DefaultQuery("language", "en"))
	country_code := strings.ToUpper(c.Query("country"))

	if country_code == "" {
		country_code = findCountryCode(c)
	}

	if isCountryExist(&country_code, db) == false {
		c.JSON(404, gin.H{
			"message": "This country doesn't exist",
		})
		return
	}

	if isLanguageExist(&language_code, db) == false {
		c.JSON(404, gin.H{
			"message": "This language doesn't exist",
		})
		return
	}

	queryStmt, err := db.Prepare("SELECT * FROM numbers_by_country_and_language($1, $2)")
	if err != nil {
		panic(err)
	}

	phonenumbers := []SQLPhonenumber{}
	rows, err := queryStmt.Query(country_code, language_code)
	defer rows.Close()

	for rows.Next() {
		var data SQLPhonenumber
		err := sqlstruct.Scan(&data, rows)

		if err == nil {
			phonenumbers = append(phonenumbers, data)
		}
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"country":      country_code,
		"language":     language_code,
		"phonenumbers": phonenumbers,
	})
}

func findCountryCode(c *gin.Context) string {
	latString := c.Query("lat")
	longString := c.Query("long")

	if latString != "" && longString != "" {
		lat, err := strconv.ParseFloat(latString, 64)
		if err != nil {
			panic(err)
		}

		long, err := strconv.ParseFloat(longString, 64)
		if err != nil {
			panic(err)
		}

		geo := c.MustGet("geo").(*georeverse.CountryReverser)
		return strings.ToUpper(geo.GetCountryCode(long, lat))
	} else {
		return ""
	}
}
