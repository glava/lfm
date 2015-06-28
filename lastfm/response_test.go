package lastfm

import "testing"
import "os"
import "io/ioutil"

func TestParsingOfArtists(t *testing.T) {
	q := ToArtists(readFile("../test-rig/top-artists.json", t))
	assert("Bee Gees", q.Topartists.Artists[0].Name, t)
}

func TestParsingOfTracks(t *testing.T) {
	q := ToTracks(readFile("../test-rig/top-tracks.json", t))
	assert("Sultans of Swing", q.Toptracks.Tracks[0].Name, t)
}

func TestPostParse(t *testing.T) {
	q := ToPlaylist(readFile("../test-rig/playlist-save.xml", t))
	assert(q.Playlists.PlaylistValue.Id, "11826164", t)

}

func assert(expected string, actual string, t *testing.T) {
	if expected != actual {
		t.Errorf("Fail. Actual %s, expected %s", actual, expected)
	}
}

func readFile(fileName string, t *testing.T) []byte {
	file, err := os.Open(fileName)
	if err != nil {
		t.Errorf("Failed to open %s file", fileName)
	}
	defer file.Close()

	b, _ := ioutil.ReadAll(file)
	return b
}
