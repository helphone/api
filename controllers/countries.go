package controllers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func GetCountries(c *gin.Context) {
	db := c.MustGet("database").(*sql.DB)
	queryStmt, err := db.Prepare("SELECT code FROM countries")
	if err != nil {
		panic(err)
	}

	var countries []string
	rows, err := queryStmt.Query()
	defer rows.Close()

	for rows.Next() {
		var code string

		if err := rows.Scan(&code); err != nil {
			panic(err)
		}
		countries = append(countries, code)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"countries": countries,
	})
}
