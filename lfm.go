package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"flag"
)

type LastFmResponse struct {
	Toptracks struct {
				  Track []struct {
					  Name string `json:"name"`
					  Duration string `json:"duration"`
					  Playcount string `json:"playcount"`
					  Listeners string `json:"listeners"`
					  Mbid string `json:"mbid"`
					  URL string `json:"url"`
					  Streamable struct {
							   Text string `json:"#text"`
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
							Artist string `json:"artist"`
							Page string `json:"page"`
							Perpage string `json:"perPage"`
							Totalpages string `json:"totalPages"`
							Total string `json:"total"`
						} `json:"@attr"`
			  } `json:"toptracks"`
}

type ApiConfig struct {
	ApiKey string `json:"api_key"`
}


func resolveUrl(apiKey string, artist string, limit int) string {
	return fmt.Sprintf("http://ws.audioscrobbler.com/2.0/?method=artist.gettoptracks&artist=%s&api_key=%s&limit=%d&format=json", artist, apiKey, limit)
}

func readApiConfig() ApiConfig {
	data, err := ioutil.ReadFile("/Users/ojkic/.lfm/config")
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

func main() {
	apiConfig := readApiConfig()

	artistName := flag.String("a", "", "artist name")
	limit := flag.Int("l", 20, "limit")
	flag.Parse()

	var url = resolveUrl(apiConfig.ApiKey, *artistName, *limit)
	body, error := executeRequest(url)

	var reponse LastFmResponse
	parseJson(body, &reponse)

	if error != nil {
		fmt.Printf("%s", error)
	} else {
		fmt.Println(reponse.Toptracks.Track[1].Name)
	}

}

