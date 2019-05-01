package services

import (
	log "github.com/sirupsen/logrus"

	"github.com/kelseyhightower/envconfig"

	"github.com/biancarosa/music-data/interfaces"
	"github.com/biancarosa/music-data/models"
)

type spotifyService struct {
	Conf struct{}
}

func (s spotifyService) GetSongInfo(name, artist string) (*models.SpotifyTrack, error) {
	return &models.SpotifyTrack{}, nil
}

func (s *spotifyService) LoadConf() {
	log.Debug("Loading configuration for SpotifyService")
	err := envconfig.Process("spotify", &s.Conf)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Debug("Env for SpotifyService loaded sucessfully")
}

//NewSpotifyService returns a implementation of a SpotifyService
func NewSpotifyService() interfaces.SpotifyService {
	var s spotifyService
	s.LoadConf()
	return &s
}
