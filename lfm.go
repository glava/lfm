package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"net/http"
	"strings"
)

type LastFmResponse struct {
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

type ApiConfig struct {
	ApiKey string `json:"api_key"`
}

//TODO: Fold this two functions in one
func resolveUrl(apiKey string, artist string, limit int) string {
	return fmt.Sprintf("http://ws.audioscrobbler.com/2.0/?method=artist.gettoptracks&artist=%s&api_key=%s&limit=%d&format=json", artist, apiKey, limit)
}

func resolveUserUrl(apiKey string, user string, limit int, period string) string {
	return fmt.Sprintf("http://ws.audioscrobbler.com/2.0/?method=user.gettoptracks&user=%s&api_key=%s&limit=%d&period=%s&format=json", user, apiKey, limit, period)
}

//TODO: solve this stupid read from home issue. currently it doesnt get the concept of ~
func readApiConfig() ApiConfig {
	data, err := ioutil.ReadFile("/Users/goranojkic/.lfm/config")
	var auth ApiConfig

	if err != nil {
		fmt.Printf("%s", err)
	}

	parseJson(data, &auth)
	return auth
}

func executeRequest(url string) (body []byte, err error) {
	response, err := http.Get(url)

	if err != nil {
		return body, err
	} else {
		defer response.Body.Close()
		var ioError error
		body, ioError = ioutil.ReadAll(response.Body)
		if ioError != nil {
			return body, ioError
		}
	}
	return body, err
}

func parseJson(jsonBody []byte, result interface{}) error {
	return json.Unmarshal(jsonBody, &result)
}

//TODO: fold to one function artistSongs and userSongs
func artistSongs(apiKey string, artist string, limit int) {
	var url = resolveUrl(apiKey, artist, limit)
	body, error := executeRequest(url)

	var response LastFmResponse
	parseJson(body, &response)

	if error != nil {
		fmt.Printf("%s", error)
	} else {

		for _, v := range response.Toptracks.Track {
			fmt.Println(artist + "-" + v.Name)
		}
	}
}

func userSongs(apiKey string, user string, limit int, period string) chan LastFmResponse {
	yield := make(chan LastFmResponse)
	go func() {
		var url = resolveUserUrl(apiKey, user, limit, period)
		body, error := executeRequest(url)

		var response LastFmResponse
		parseJson(body, &response)

		if error != nil {
			fmt.Printf("%s", error)
		} else {

			yield <- response
		}
	}()

	return yield
}

func outputTracks(response LastFmResponse, optionalTitle string) {
	if optionalTitle != "" {
		color.Green(optionalTitle)
	}
	for _, v := range response.Toptracks.Track {
		color.Yellow("\t" + v.Artist.Name + " - " + v.Name)
	}
}

func main() {
	apiConfig := readApiConfig()
	artistName := flag.String("a", "", "artist name")
	limit := flag.Int("l", 20, "limit")

	userNames := flag.String("u", "", "last.fm user names comma separated")
	period := flag.String("p", "", "period you are searching for")

	flag.Parse()

	if *artistName != "" {
		artistSongs(apiConfig.ApiKey, *artistName, *limit)
	} else {
		for _, user := range strings.Split(*userNames, ",") {
			outputTracks(<-userSongs(apiConfig.ApiKey, user, *limit, *period), user)
		}
	}

}
