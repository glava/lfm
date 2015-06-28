package lastfm

import "testing"
import "crypto/md5"
import "encoding/hex"

func TestSignature(t *testing.T) {
	type params map[string]string
	res := Signature(params{"x": "3", "a": "1", "b": "2"}, "secret")

	hasher := md5.New()
	hasher.Write([]byte("a1b2x3" + "secret"))

	if res != hex.EncodeToString(hasher.Sum(nil)) {
		t.Errorf("dont work %s", res)
	}
}
