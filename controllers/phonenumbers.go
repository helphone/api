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
		if c.Query("lat") != "" && c.Query("long") != "" {
			country_code = findCountryCodeFromLatAndLong(c)
		} else {
			country_code = "*"
		}
	} else {
		if isCountryExist(&country_code, db) == false {
			c.JSON(404, gin.H{
				"message": "This country doesn't exist",
			})
			return
		}
	}

	if isLanguageExist(&language_code, db) == false {
		c.JSON(404, gin.H{
			"message": "This language doesn't exist",
		})
		return
	}

	var rows *sql.Rows

	if country_code == "*" {
		rows = getPhonenumbers(db, language_code)
	} else {
		rows = getPhonenumbersForCountry(db, country_code, language_code)
	}

	defer rows.Close()
	countries := []*Country{}

	for rows.Next() {
		var data Phonenumber
		err := sqlstruct.Scan(&data, rows)

		if err != nil {
			panic(err)
		}

		var countryPosition int = -1
		for i := 0; i < len(countries); i++ {
			if countries[i].Code == data.Country {
				countryPosition = i
			}
		}

		if countryPosition != -1 {
			countries[countryPosition].Phonenumbers = append(countries[countryPosition].Phonenumbers, &data)
		} else {
			new_country := Country{
				data.Country,
				[]*Phonenumber{},
			}
			countries = append(countries, &new_country)
		}
	}

	err := rows.Err()
	if err != nil {
		panic(err)
	}

	response := JSONResponse{language_code, countries}
	c.JSON(200, response)

}

func getPhonenumbers(db *sql.DB, language_code string) *sql.Rows {
	queryStmt, err := db.Prepare("SELECT * FROM numbers_by_language($1)")
	if err != nil {
		panic(err)
	}

	rows, err := queryStmt.Query(language_code)
	return rows
}

func getPhonenumbersForCountry(db *sql.DB, country_code string, language_code string) *sql.Rows {
	queryStmt, err := db.Prepare("SELECT * FROM numbers_by_country_and_language($1, $2)")
	if err != nil {
		panic(err)
	}

	rows, err := queryStmt.Query(country_code, language_code)
	return rows
}

func findCountryCodeFromLatAndLong(c *gin.Context) string {
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
