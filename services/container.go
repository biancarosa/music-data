package services

import (
	"github.com/biancarosa/music-data/interfaces"
)

//Container is the container for services
type Container struct {
	SongService interfaces.SongService
}

//Register registers services on a container
func Register() *Container {
	return &Container{
		SongService: NewSongService(),
	}
}
