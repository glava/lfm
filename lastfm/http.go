package lastfm

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func ArtistUrl(apiKey string, artist string, limit int) string {
	return fmt.Sprintf("http://ws.audioscrobbler.com/2.0/?method=artist.gettoptracks&artist=%s&api_key=%s&limit=%d&format=json", cleanParam(artist), apiKey, limit)
}

func UserUrl(apiKey string, user string, limit int, period string) string {
	return fmt.Sprintf("http://ws.audioscrobbler.com/2.0/?method=user.gettoptracks&user=%s&api_key=%s&limit=%d&period=%s&format=json", user, apiKey, limit, period)
}

func TagUrl(apiKey string, tag string, limit int) string {
	return fmt.Sprintf("http://ws.audioscrobbler.com/2.0/?method=tag.gettopartists&tag=%s&api_key=%s&limit=%d&format=json", tag, apiKey, limit)
}

func TokenUrl(apiKey string, apiSig string) string {
	return fmt.Sprintf("http://ws.audioscrobbler.com/2.0/?method=auth.gettoken&api_key=%s&api_sig=%s&format=json", apiKey, apiSig)
}

func Execute(url string) chan []byte {
	yield := make(chan []byte)
	go func() {
		body, error := getHttpBody(url)
		if error != nil {
			fmt.Printf("%s", error)
		} else {
			yield <- body
		}
	}()

	return yield
}

func cleanParam(param string) string {
	return strings.Replace(param, " ", "+", -1)
}

func getHttpBody(url string) (body []byte, err error) {
	response, err := http.Get(url)

	if err != nil {
		return body, err
	} else {
		defer response.Body.Close()
		var ioError error
		body, ioError = ioutil.ReadAll(response.Body)

		if ioError != nil {
			return body, ioError
		}
	}
	return body, err
}
