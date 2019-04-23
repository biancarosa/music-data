package routes

import (
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"

	"github.com/biancarosa/music-data/models"
)

//HealthCheck is the route that prints a sucessful message when the application is fine.
func HealthCheck(c echo.Context) error {
	log.Debug("Healthcheck has been called")
	msg := "I seem to be perfectly fine"
	apiResponse := &models.APIResponse{
		Message: &msg,
	}
	return c.JSON(http.StatusOK, apiResponse)
}
