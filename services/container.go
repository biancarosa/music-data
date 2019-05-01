package services

import (
	log "github.com/sirupsen/logrus"

	"github.com/biancarosa/music-data/interfaces"
)

//Container is the container for services
type Container struct {
	SongService interfaces.SongService
}

//Register registers services on a container
func Register() *Container {
	log.Debug("Initializing containers")
	lastFMService := NewLastFMService()
	spotifyService := NewSpotifyService()
	return &Container{
		SongService: NewSongService(lastFMService, spotifyService),
	}
}
