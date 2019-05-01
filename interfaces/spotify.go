package interfaces

import (
	"github.com/biancarosa/music-data/models"
)

//SpotifyService is the interface that defines a SpotifyService provider
type SpotifyService interface {
	GetSongInfo(name, artist string) (*models.SpotifyTrack, error)
}
