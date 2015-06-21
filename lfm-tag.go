package main

import (
	"./config"
	"./helper"
	"./lastfm"
	"flag"
	"strings"
)

func main() {
	apiConfig := config.Load()

	limit := flag.Int("l", 20, "limit")
	tags := flag.String("t", "", "last.fm tags")

	flag.Parse()

	for _, tag := range strings.Split(*tags, ",") {
		helper.Output(lastfm.ToArtists(<-lastfm.Execute(lastfm.TagUrl(apiConfig.ApiKey, tag, *limit))), "")
	}

}
