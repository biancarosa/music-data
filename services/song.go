package services

import (
	"github.com/biancarosa/music-data/interfaces"
	"github.com/biancarosa/music-data/models"
)

type songService struct {
	lastFMService interfaces.LastFMService
}

func (s songService) GetSongInfo(name, artist string) (*models.Song, error) {
	return s.lastFMService.GetSongInfo(name, artist)
}

//NewSongService returns a implementation of a SongService
func NewSongService(lastFMService interfaces.LastFMService) interfaces.SongService {
	return &songService{lastFMService: lastFMService}
}
