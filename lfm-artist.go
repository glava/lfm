package main

import (
	"./config"
	"./lastfm"
	"flag"
	"./helper"
	"strings"
)

func main() {
	apiConfig := config.Load()

	artistsFlag := flag.String("a", "", "artists names  - separated by coma. Example: -a Drake,Kanye West")
	limit := flag.Int("l", 20, "number of songs you want from artist. Default 20")
	flag.Parse()

	artists := append(strings.Split(*artistsFlag, ","), helper.ReadStdin()...)

	for _, artist := range artists {
		if len(strings.TrimSpace(artist)) > 0 {
			helper.Output(<-lastfm.Execute(lastfm.ArtistUrl(apiConfig.ApiKey, artist, *limit)), "")
		}
	}

}
