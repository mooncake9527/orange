package rsaUtil

import (
	"github.com/mooncake9527/orange/common/consts"
	"testing"
)

func TestDecrypt(t *testing.T) {
	msg := "hello world"
	d, _ := Encrypt(msg, consts.PubKey)
	d1, _ := Decrypt(d, consts.PriKey)
	if d1 != msg {
		t.Error("decrypt error")
	}
}
