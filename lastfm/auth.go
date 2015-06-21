package lastfm

import ".././config"
import "crypto/md5"
import "encoding/hex"

func FetchRequestToken(config config.ApiConfig) string {
	hasher := md5.New()
	hasher.Write([]byte("api_key" + config.ApiKey + "methodauth.getSessiontoken" + config.ApiSecret))
	api_sig := hex.EncodeToString(hasher.Sum(nil))
	return ToToken(<-Execute(TokenUrl(config.ApiKey, api_sig))).ToString()
}
