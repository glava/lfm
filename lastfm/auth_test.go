package lastfm

import "testing"
import "github.com/glava/lfm/lastfm"
import "github.com/glava/lfm/config"
import "crypto/md5"
import "encoding/hex"
import "fmt"

func TestSignature(t *testing.T) {
	type params map[string]string
	res := lastfm.Signature(params{"x": "3", "a": "1", "b": "2"}, "secret")

	hasher := md5.New()
	hasher.Write([]byte("a1b2x3" + "secret"))

	if res != hex.EncodeToString(hasher.Sum(nil)) {
		t.Errorf("dont work %s", res)
	}
}

func TestPost(t *testing.T) {
	apiConfig := config.Load()
	session := config.GetSession()
	fmt.Println(session)
	params := map[string]string{"method": "playlist.create", "title": "new one", "api_key": apiConfig.ApiKey, "sk": session}
	sig := lastfm.Signature(params, apiConfig.ApiSecret)
	params["api_sig"] = sig
	fmt.Println(params["api_sig"])
	fmt.Println(<-lastfm.Post(params))

	t.Errorf("dont work")
}
