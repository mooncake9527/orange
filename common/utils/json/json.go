package jsonUtil

import (
	"encoding/json"
	"github.com/mooncake9527/x/xerrors/xerror"
)

func ToJsonString(v any) string {
	bytes, _ := json.Marshal(v)
	return string(bytes)
}

func Marshal(v any) ([]byte, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return nil, xerror.Wrap(err, "json marshal")
	}
	return bytes, nil
}

func Unmarshal(data []byte, v any) error {
	err := json.Unmarshal(data, v)
	if err != nil {
		return xerror.Wrap(err, "json unmarshal")
	}
	return nil
}
