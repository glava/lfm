package main

import (
	"./config"
	"./lastfm"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"strings"
)

func output(response lastfm.ArtistResponse, optionalTitle string) {
	if optionalTitle != "" {
		color.Green(optionalTitle)
	}
	for _, v := range response.Topartists.Artist {
		fmt.Println(v.Name)
	}
}

func main() {
	apiConfig := config.Load()

	limit := flag.Int("l", 20, "limit")
	tags := flag.String("t", "", "last.fm tags")

	flag.Parse()

	for _, tag := range strings.Split(*tags, ",") {
		output(<-lastfm.ArtistExecute(lastfm.TagUrl(apiConfig.ApiKey, tag, *limit)), "")
	}

}
