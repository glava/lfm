package lastfm

import "github.com/glava/lfm/config"
import "crypto/md5"
import "encoding/hex"
import "fmt"
import "os"
import "bufio"
import "github.com/skratchdot/open-golang/open"
import "sort"

func FetchRequestToken(config config.ApiConfig) string {
	params := map[string]string{"api_key": config.ApiKey, "method": "auth.gettoken"}
	return ToToken(<-Get(TokenUrl(config.ApiKey, Signature(params, config.ApiSecret)))).ToString()
}

func FetchSession(config config.ApiConfig, token string) string {

	params := map[string]string{"api_key": config.ApiKey, "token": token, "method": "auth.getsession"}

	open.Run(fmt.Sprintf("http://www.last.fm/api/auth/?api_key=%s&token=%s", config.ApiKey, token))
	fmt.Println("After you allow access app to your last.fm profile press enter")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
	return ToSession(<-Get(SessionUrl(config.ApiKey, token, Signature(params, config.ApiSecret)))).Session.Key
}

func Signature(params map[string]string, apiSecret string) string {
	sig := new(string)
	sortedKeys := make([]string, len(params))
	i := 0
	for k := range params {
		sortedKeys[i] = k
		i++
	}
	sort.Strings(sortedKeys)

	for _, v := range sortedKeys {
		*sig = *sig + v + params[v]
	}

	hasher := md5.New()
	hasher.Write([]byte(*sig + apiSecret))
	return hex.EncodeToString(hasher.Sum(nil))

}
