package services

import (
	"github.com/biancarosa/music-data/interfaces"
	"github.com/biancarosa/music-data/models"
)

type songService struct {
}

func (s songService) GetSongInfo(name, artist string) (*models.Song, error) {
	return &models.Song{
		Name:   name,
		Artist: artist,
	}, nil
}

//NewSongService returns a implementation of a SongService
func NewSongService() interfaces.SongService {
	return &songService{}
}
