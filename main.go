package main

import (
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/helpnumber/api/config"
	c "github.com/helpnumber/api/controllers"
)

func main() {
	r := gin.Default()
	r.Use(config.Provide())
	r.Use(config.ETag())
	r.Use(cors.Default())

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Ping")
	})
	r.GET("/countries", c.GetCountries)
	r.GET("/phonenumbers", c.GetPhonenumbers)

	r.Run(":3000")
}
