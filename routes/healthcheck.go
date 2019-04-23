package routes

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

type applicationDetails struct {
	// Número de go rotinas abertas
	//
	// Required: true
	Goroutines int `json:"goroutines"`
	// Quanto de memória está alocada para uso da aplicação
	//
	// Required: true
	HeapAlloc string `json:"heapAlloc"`
}

// swagger:response healthCheckResponse
type healthCheckResponse struct {
	// Versão da aplicação
	//
	// Required: true
	ApplicationVersion string             `json:"applicationVersion"`
	Application        applicationDetails `json:"application"`
}

// HealthCheck is the route that prints a sucessful message when the application is fine.
// swagger:route GET /healthcheck HealthCheck
// Retorna informação da saúde da aplicação
// responses:
//  200: healthCheckResponse
func HealthCheck(c echo.Context) error {
	log.Debug("Healthcheck has been called")
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	apiResponse := &healthCheckResponse{
		ApplicationVersion: "0.1.0", // #TODO Get from conf file
		Application: applicationDetails{
			Goroutines: runtime.NumGoroutine(),
			HeapAlloc:  fmt.Sprintf("%d mb", mem.Sys/1000000), // bytes
		},
	}
	return c.JSON(http.StatusOK, apiResponse)
}
