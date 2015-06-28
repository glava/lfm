package lastfm

type Track struct {
	Name       string `json:"name"`
	Duration   string `json:"duration"`
	Playcount  string `json:"playcount"`
	Listeners  string `json:"listeners"`
	Mbid       string `json:"mbid"`
	URL        string `json:"url"`
	Streamable struct {
		Text      string `json:"#text"`
		Fulltrack string `json:"fulltrack"`
	} `json:"streamable"`
	Artist struct {
		Name string `json:"name"`
	} `json:"artist"`
	Image []struct {
		Text string `json:"#text"`
		Size string `json:"size"`
	} `json:"image"`
	Attr struct {
		Rank string `json:"rank"`
	} `json:"@attr"`
}

type Artist struct {
	Name       string `json:"name"`
	Mbid       string `json:"mbid"`
	URL        string `json:"url"`
	Streamable string `json:"streamable"`
	Image      []struct {
		Text string `json:"#text"`
		Size string `json:"size"`
	} `json:"image"`
	Attr struct {
		Rank string `json:"rank"`
	} `json:"@attr"`
}

type Playlist struct {
	Id    string `xml:"id"`
	Title string `xml:"title"`
	Url   string `xml:"url"`
}
