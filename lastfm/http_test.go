package lastfm

import (
	"fmt"
	"github.com/glava/lfm/config"
	"github.com/glava/lfm/lastfm"
	"io/ioutil"
	"os"
	"testing"
)

func TestPostSignature(t *testing.T) {
	apiConfig := new(config.ApiConfig)
	apiConfig.ApiKey = "062928747ddd9617d0bac5c36286b980"
	apiConfig.ApiSecret = "5fde6ddefa7bb830e093cfa4082d6b9c"
	session := "3b77cbd9843e526d86ffb1abf994959c"

	params := map[string]string{"method": "playlist.create", "title": "fresh one2", "api_key": apiConfig.ApiKey, "sk": session}

	sig := lastfm.Signature(params, apiConfig.ApiSecret)
	expectedSig := "8c02305b691ed553558c47d7539cc11c"

	if sig != expectedSig {
		t.Errorf("Failed to generate proper signature. Got %s ,expected %s", sig, expectedSig)
	}
}

func TestPostParse(t *testing.T) {
	xmlFile, err := os.Open("../test-rig/playlist-save.xml")
	if err != nil {
		t.Errorf("Failed to open playlist-save.xml file")
	}

	defer xmlFile.Close()

	b, _ := ioutil.ReadAll(xmlFile)
	q := lastfm.ToPlaylist(b)
	if q.Playlists.Playlist.Id != "11826164" {
		t.Errorf("Failed to parse. Id of parse %s is not equal to expected %s", q.Playlists.Playlist.Id, "11826164")
	}

}
