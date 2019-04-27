package services

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"

	"github.com/kelseyhightower/envconfig"

	"github.com/biancarosa/music-data/interfaces"
	"github.com/biancarosa/music-data/models"
)

type lastFMService struct {
	Conf struct {
		Token string `required:"true"`
	}
}

//LastFMResponse defines the response of a last fm track
type LastFMResponse struct {
	Track struct {
		Name       string `json:"name"`
		Mbid       string `json:"mbid"`
		URL        string `json:"url"`
		Duration   string `json:"duration"`
		Streamable struct {
			Text      string `json:"#text"`
			Fulltrack string `json:"fulltrack"`
		} `json:"streamable"`
		Listeners string `json:"listeners"`
		Playcount string `json:"playcount"`
		Artist    struct {
			Name string `json:"name"`
			Mbid string `json:"mbid"`
			URL  string `json:"url"`
		} `json:"artist"`
		Album struct {
			Artist string `json:"artist"`
			Title  string `json:"title"`
			Mbid   string `json:"mbid"`
			URL    string `json:"url"`
			Image  []struct {
				Text string `json:"#text"`
				Size string `json:"size"`
			} `json:"image"`
			Attr struct {
				Position string `json:"position"`
			} `json:"@attr"`
		} `json:"album"`
		Toptags struct {
			Tag []struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"tag"`
		} `json:"toptags"`
		Wiki struct {
			Published string `json:"published"`
			Summary   string `json:"summary"`
			Content   string `json:"content"`
		} `json:"wiki"`
	} `json:"track"`
}

func (s lastFMService) GetURL() string {
	return fmt.Sprintf("http://ws.audioscrobbler.com/2.0/?api_key=%s&format=json", s.Conf.Token)
}

func (s lastFMService) GetSongInfo(name, artist string) (*models.Song, error) {
	url := fmt.Sprintf("%s&method=track.getInfo&artist=%s&track=%s", s.GetURL(), artist, name)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Debugf("Returned body %s", body)
	var lastFMResponse LastFMResponse
	err = json.Unmarshal(body, &lastFMResponse)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Debugf("Response is %#v", lastFMResponse)
	return &models.Song{
		Name:   lastFMResponse.Track.Name,
		Artist: lastFMResponse.Track.Artist.Name,
	}, nil
}

func (s *lastFMService) LoadConf() {
	log.Debug("Loading configuration for LastFMService")
	err := envconfig.Process("lastFM", &s.Conf)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Debug("Env for LastFMService loaded sucessfully")
}

//NewLastFMService returns a implementation of a LastFMService
func NewLastFMService() interfaces.LastFMService {
	var s lastFMService
	s.LoadConf()
	return &s
}
