package lastfm

import "testing"
import "github.com/glava/lfm/lastfm"
import "github.com/glava/lfm/config"
import "fmt"

func TestPost(t *testing.T) {
	apiConfig := config.Load()
	session := config.GetSession()
	fmt.Println(session)
	params := map[string]string{"method": "playlist.create", "title": "fresh one", "api_key": apiConfig.ApiKey, "sk": session}
	sig := lastfm.Signature(params, apiConfig.ApiSecret)
	params["api_sig"] = sig
	fmt.Println(params["api_sig"])
	fmt.Println(<-lastfm.Post(params))
}
