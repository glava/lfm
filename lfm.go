package main

import (
	"./config"
	"./lastfm"
	"flag"
	"github.com/fatih/color"
	"strings"
)

func output(response lastfm.Response, optionalTitle string) {
	if optionalTitle != "" {
		color.Green(optionalTitle)
	}
	for _, v := range response.Toptracks.Track {
		color.Yellow("\t" + v.Artist.Name + " - " + v.Name)
	}
}

func main() {
	apiConfig := config.Load()

	artistName := flag.String("a", "", "artist name")
	limit := flag.Int("l", 20, "limit")
	userNames := flag.String("u", "", "last.fm user names comma separated")
	period := flag.String("p", "", "period you are searching for")

	flag.Parse()

	if *artistName != "" {
		output(<-lastfm.Execute(lastfm.ArtistUrl(apiConfig.ApiKey, *artistName, *limit)), "")
	} else {
		for _, user := range strings.Split(*userNames, ",") {
			output(<-lastfm.Execute(lastfm.UserUrl(apiConfig.ApiKey, user, *limit, *period)), user)
		}
	}

}
