package interfaces

import (
	"github.com/biancarosa/music-data/models"
)

//LastFMService is the interface that defines a LastFMService provider
type LastFMService interface {
	GetSongInfo(name, artist string) (*models.Song, error)
}
