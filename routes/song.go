package routes

import (
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

// songResponse defines the response that you get of a valid song.
// swagger:response songResponse
type songResponse struct {
	// The response for a valid song request.
	//
	// in: body
	Body struct {
		// Song name
		//
		// Required: true
		Name string `json:"name"`
	}
}

// Song is the route that receives a song name and returns song info.
func Song(c echo.Context) error {
	// swagger:route GET /song Song
	//
	// Returns information about a song.
	//
	// We will retrieve some info in other APIs and in our database.
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: https
	//
	//     Responses:
	//     200: songResponse
	log.Debug("GetSong has been called")
	name := c.QueryParam("name")
	if name == "" {
		resp := validationError{
			queryParameters: []string{"name"},
		}.BuildResponse()
		return c.JSON(http.StatusBadRequest, resp.Body) // #TODO Add this in docs
	}
	var resp songResponse
	resp.Body.Name = name
	return c.JSON(http.StatusOK, resp.Body)
}
