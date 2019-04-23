package routes

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

// songResponse defines the response that you get of a valid healthcheck.
// swagger:response healthCheckResponse
type healthCheckResponse struct {
	// The response for a valid healthcheck request.
	//
	// in: body
	Body struct {
		// Application version
		//
		// Required: true
		ApplicationVersion string `json:"applicationVersion"`
		Application        struct {
			// Número de go rotinas abertas
			//
			// Required: true
			Goroutines int `json:"goroutines"`
			// Quanto de memória está alocada para uso da aplicação
			//
			// Required: true
			HeapAlloc string `json:"heapAlloc"`
		}
	}
}

// HealthCheck is the route that prints a sucessful message when the application is fine.
func HealthCheck(c echo.Context) error {
	// swagger:route GET /healthcheck HealthCheck
	//
	// Returns information about our application health.
	//
	// We will check some info regarding the application health, including memory stats.
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: https
	//
	//     Responses:
	//     200: healthCheckResponse
	log.Debug("Healthcheck has been called")
	var mem runtime.MemStats
	var resp healthCheckResponse

	runtime.ReadMemStats(&mem)

	resp.Body.ApplicationVersion = "0.1.0" // #TODO Get from conf file
	resp.Body.Application.Goroutines = runtime.NumGoroutine()
	resp.Body.Application.HeapAlloc = fmt.Sprintf("%d mb", mem.Sys/1000000) // bytes

	return c.JSON(http.StatusOK, resp.Body)
}
