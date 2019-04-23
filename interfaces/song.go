package interfaces

import (
	"github.com/biancarosa/music-data/models"
)

//SongService is the interface that defines a SongService provider
type SongService interface {
	GetSongInfo(name, artist string) (*models.Song, error)
}
