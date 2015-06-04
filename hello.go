package main

import (
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
		"encoding/json"
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

func main() {

	response, err := http.Get("http://ws.audioscrobbler.com/2.0/?method=artist.gettoptracks&artist=Drake&api_key=5feb650940181e81c0cbcb09f8e58c03&limit=20&format=json")
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		rawContent, err := ioutil.ReadAll(response.Body)
		var result LastFmResponse

		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}

		jsonErr := json.Unmarshal(rawContent, &result)

		if jsonErr != nil {
			fmt.Printf("%s", jsonErr)
			os.Exit(1)
		}

		fmt.Println(result.Toptracks.Track[1].Name)

	}
}

