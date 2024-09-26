package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/mooncake9527/orange/common/utils/whiteip"
	common "github.com/mooncake9527/x/xutil/ip"
)

// 白名单 权限检查中间件
func WhileIp() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := common.GetClientIP(c)
		if !whiteip.InWhile(ip) {
			Fail(c, 500, "请联系管理员添加白名单")
			return
		}
		c.Next()
	}
}
