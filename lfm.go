package main

import (
	"flag"
	"github.com/glava/lfm/config"
	"github.com/glava/lfm/helper"
	"github.com/glava/lfm/lastfm"
	"strings"
)

func main() {
	apiConfig := config.Load()

	artistsFlag := flag.String("a", "", "artists names  - separated by coma. Example: -a Drake,Kanye West")
	userNames := flag.String("u", "", "last.fm user names comma separated")
	limit := flag.Int("l", 20, "number of songs you want from artist. Default 20")
	period := flag.String("p", "", "period you are searching for")
	tags := flag.String("t", "", "last.fm tags")
	playlist := flag.String("pl", "", "playlist")

	flag.Parse()

	//config.SaveToken(lastfm.FetchRequestToken(apiConfig))
	//config.SaveSession(lastfm.FetchSession(apiConfig, config.GetToken()))

	artists := append(strings.Split(*artistsFlag, ","), helper.ReadStdin()...)

	//TODO: move go

	if *artistsFlag != "" {

		for _, artist := range artists {
			if len(strings.TrimSpace(artist)) > 0 {
				helper.Output(lastfm.ToTracks(<-lastfm.Get(lastfm.ArtistUrl(apiConfig.ApiKey, artist, *limit))), "")
			}
		}
	}

	if *userNames != "" {
		for _, user := range strings.Split(*userNames, ",") {
			helper.Output(lastfm.ToTracks(<-lastfm.Get(lastfm.UserUrl(apiConfig.ApiKey, user, *limit, *period))), user)
		}
	}

	if *tags != "" {
		for _, tag := range strings.Split(*tags, ",") {
			helper.Output(lastfm.ToArtists(<-lastfm.Get(lastfm.TagUrl(apiConfig.ApiKey, tag, *limit))), "")
		}
	}

	if *playlist != "" {
		session := config.GetSession()
		params := map[string]string{"method": "playlist.create", "title": *playlist, "api_key": apiConfig.ApiKey, "sk": session}
		sig := lastfm.Signature(params, apiConfig.ApiSecret)
		params["api_sig"] = sig
		helper.Output(lastfm.ToPlaylist(<-lastfm.Post(params)), "")
	}
}
