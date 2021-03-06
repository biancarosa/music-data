// Package main Music Data API.
//
// the purpose of this application is to provide information about music.
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http, https
//     Host: localhost
//     BasePath: /v1
//     Version: 0.1.0
//     License: MIT http://opensource.org/licenses/MIT
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"github.com/biancarosa/music-data/services"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"

	"github.com/biancarosa/music-data/routes"
)

func main() {
	e := echo.New()

	// Setup Logrus
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	container := services.Register()

	v1 := e.Group("/v1")
	v1.GET("/healthcheck", routes.HealthCheck)
	v1.GET("/song", func(c echo.Context) error {
		return routes.Song(c, container.SongService)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
