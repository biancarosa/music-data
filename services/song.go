package services

import (
	"github.com/biancarosa/music-data/interfaces"
	"github.com/biancarosa/music-data/models"
)

type songService struct {
	lastFMService  interfaces.LastFMService
	spotifyService interfaces.SpotifyService
}

func (s songService) GetSongInfo(name, artist string) (*models.Song, error) {
	lastFMTrack, err := s.lastFMService.GetSongInfo(name, artist)
	if err != nil {
		return nil, err
	}
	_, err = s.spotifyService.GetSongInfo(name, artist)
	if err != nil {
		return nil, err
	}
	return &models.Song{
		Artist: lastFMTrack.Artist.Name,
		Name:   lastFMTrack.Name,
	}, nil
}

//NewSongService returns a implementation of a SongService
func NewSongService(lastFMService interfaces.LastFMService, spotifyService interfaces.SpotifyService) interfaces.SongService {
	return &songService{lastFMService: lastFMService, spotifyService: spotifyService}
}
