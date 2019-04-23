package routes

import (
	"github.com/biancarosa/music-data/interfaces"
	"github.com/biancarosa/music-data/models"
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
	Body *models.Song
}

// Song is the route that receives a song name and returns song info.
func Song(c echo.Context, songService interfaces.SongService) error {
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
	artist := c.QueryParam("artist")
	if name == "" || artist == "" {
		resp := validationError{
			queryParameters: []string{},
		}
		if name == "" {
			resp.queryParameters = append(resp.queryParameters, "name")
		}
		if artist == "" {
			resp.queryParameters = append(resp.queryParameters, "artist")
		}
		return c.JSON(http.StatusBadRequest, resp.BuildResponse().Body) // #TODO Add this in docs
	}
	var resp songResponse
	var err error
	resp.Body, err = songService.GetSongInfo(name, artist)
	if err != nil {
		var errorResp errorResponse
		errorResp.Body.Message = internalServerError
		return c.JSON(http.StatusInternalServerError, errorResp)
	}
	return c.JSON(http.StatusOK, resp.Body)
}
