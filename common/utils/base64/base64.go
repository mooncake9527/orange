package base64Util

import (
	"encoding/base64"
	"github.com/mooncake9527/x/xerrors/xerror"
)

func Encode(d []byte) []byte {
	enc := base64.StdEncoding
	buf := make([]byte, enc.EncodedLen(len(d)))
	enc.Encode(buf, d)
	return buf
}

func EncodeToString(d []byte) string {
	return base64.StdEncoding.EncodeToString(d)
}

func Decode(d string) ([]byte, error) {
	bytes, err := base64.StdEncoding.DecodeString(d)
	if err != nil {
		return nil, xerror.New(err.Error())
	}
	return bytes, nil
}
