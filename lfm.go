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
	userNames := flag.String("u", "", "last.fm user names comma separated")
	period := flag.String("p", "", "period you are searching for")
	flag.Parse()

	for _, user := range strings.Split(*userNames, ",") {
		helper.Output(<-lastfm.Execute(lastfm.UserUrl(apiConfig.ApiKey, user, *limit, *period)), user)
	}

}
