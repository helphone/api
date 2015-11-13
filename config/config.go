package config

import (
	"database/sql"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/helphone/geo/georeverse"
	_ "github.com/lib/pq"
)

func Provide() gin.HandlerFunc {
	hostname := os.Getenv("DB_PORT_5432_TCP_ADDR")
	port := os.Getenv("DB_PORT_5432_TCP_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	db, err := sql.Open("postgres", "postgres://"+username+":"+password+"@"+hostname+":"+port+"/helpnumber?sslmode=disable")
	if err != nil && db.Ping() != nil {
		panic(err)
	}

	geo, err := georeverse.NewCountryReverser("./files/polygons.properties")
	if err != nil {
		panic(err)
	}

	return func(c *gin.Context) {
		c.Header("Connection", "keep-alive")
		c.Header("Cache-Control", "max-age=0, private, must-revalidate")
		c.Header("X-Content-Type-Options", "nosniff")

		c.Set("database", db)
		c.Set("geo", geo)

		c.Next()
	}
}
