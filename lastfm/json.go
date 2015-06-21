package lastfm

import "encoding/json"

type Response interface {
	ToString() string
}

type TracksResponse struct {
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

func (t TracksResponse) ToString() string {
	s := ""
	for _, v := range t.Toptracks.Track {
		s = s + v.Artist.Name + " - " + v.Name + "\n"
	}
	return s
}

type TokenResponse struct {
	Token string `json:"token"`
}

func (t TokenResponse) ToString() string {
	return t.Token
}

type ArtistResponse struct {
	Topartists struct {
		Artist []struct {
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
		} `json:"artist"`
		Attr struct {
			Tag string `json:"tag"`
		} `json:"@attr"`
	} `json:"topartists"`
}

func (t ArtistResponse) ToString() string {
	s := ""
	for _, v := range t.Topartists.Artist {
		s = s + v.Name + "\n"
	}
	return s
}

func ToTracks(httpBody []byte) TracksResponse {
	var response TracksResponse
	json.Unmarshal(httpBody, &response)
	return response
}

func ToArtists(httpBody []byte) ArtistResponse {
	var response ArtistResponse
	json.Unmarshal(httpBody, &response)
	return response
}

func ToToken(httpBody []byte) TokenResponse {
	var response TokenResponse
	json.Unmarshal(httpBody, &response)
	return response
}
