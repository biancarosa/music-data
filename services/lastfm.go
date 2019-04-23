package services

import (
	log "github.com/sirupsen/logrus"

	"github.com/kelseyhightower/envconfig"

	"github.com/biancarosa/music-data/interfaces"
	"github.com/biancarosa/music-data/models"
)

type lastFMService struct {
	Conf struct {
		Token string `required:"true"`
	}
}

func (s lastFMService) GetSongInfo(name, artist string) (*models.Song, error) {
	return &models.Song{
		Name:   name,
		Artist: artist,
	}, nil
}

func (s lastFMService) LoadConf() {
	log.Debug("Loading configuration for LastFMService")
	err := envconfig.Process("lastFM", &s.Conf)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Debugf("Env for LastFMService loaded sucessfully")
}

//NewLastFMService returns a implementation of a LastFMService
func NewLastFMService() interfaces.LastFMService {
	var s lastFMService
	s.LoadConf()
	return &s
}
