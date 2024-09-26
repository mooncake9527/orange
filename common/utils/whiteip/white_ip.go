package whiteip

import (
	"github.com/mooncake9527/orange/common/config"
	"log/slog"
	"strings"
)

func InWhile(ip string) bool {
	slog.Info("InWhile", "ip", ip)
	if ip == "127.0.0.1" || ip == "::1" {
		return true
	}
	return strings.Contains(config.Ext.WhileIps, ip)
}
