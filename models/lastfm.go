package models

//LastFMTrack defines a track from LastFM
type LastFMTrack struct {
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
}
