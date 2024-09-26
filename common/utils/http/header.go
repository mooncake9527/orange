package httpUtil

import "net/http"

var (
	appKeyKeys = []string{"AppId", "appid", "appId", "app-id", "Appid", "App-id"}
)

func GetHeader(headerKeys []string, header http.Header) string {
	for _, headerKey := range headerKeys {
		if v := header.Get(headerKey); v != "" {
			return v
		}
	}
	return ""
}

func GetAppKeyFromHeader(header http.Header) string {
	for _, headerKey := range appKeyKeys {
		if v := header.Get(headerKey); v != "" {
			return v
		}
	}
	return ""
}
