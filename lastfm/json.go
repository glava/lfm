package lastfm

type Response struct {
	Toptracks struct {
		Track []struct {
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
		} `json:"track"`
		Attr struct {
			Artist     string `json:"artist"`
			Page       string `json:"page"`
			Perpage    string `json:"perPage"`
			Totalpages string `json:"totalPages"`
			Total      string `json:"total"`
		} `json:"@attr"`
	} `json:"toptracks"`
}
