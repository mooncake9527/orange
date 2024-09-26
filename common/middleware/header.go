package middleware

import (
	gLocal "github.com/mooncake9527/orange-core/common/xlog/g_local"
	ctxUtil "github.com/mooncake9527/orange/common/utils/ctx"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mooncake9527/orange-core/common/utils"
)

// NoCache is a middleware function that appends headers
// to prevent the client from caching the HTTP response.
func NoCache(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
	c.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
	c.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
	c.Next()
}

// 获取请求id，没有默认一个
func ReqId(c *gin.Context) {
	gLocal.SetReqId(utils.GetReqId(c))
	defer gLocal.SetReqId("")
	c.Next()
}

func AppKey(c *gin.Context) {
	gLocal.SetAppKey(ctxUtil.GetAppKey(c))
	defer gLocal.SetAppKey("")
	c.Next()
}

// // Options is a middleware function that appends headers
// // for options requests and aborts then exits the middleware
// // chain and ends the request.
// func Options(c *gin.Context) {
// 	if c.Request.Method != "OPTIONS" {
// 		c.Next()
// 	} else {
// 		c.Header("Access-Control-Allow-Origin", "*")
// 		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
// 		c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
// 		c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
// 		c.Header("Content-Type", "application/json")
// 		c.AbortWithStatus(200)
// 	}
// }

// // Secure is a middleware function that appends security
// // and resource access headers.
// func Secure(c *gin.Context) {
// 	c.Header("Access-Control-Allow-Origin", "*")
// 	//c.Header("X-Frame-Options", "DENY")
// 	c.Header("X-Content-Type-Options", "nosniff")
// 	c.Header("X-XSS-Protection", "1; mode=block")
// 	if c.Request.TLS != nil {
// 		c.Header("Strict-Transport-Security", "max-age=31536000")
// 	}

// 	// Also consider adding Content-Security-Policy headers
// 	// c.Header("Content-Security-Policy", "script-src 'self' https://cdnjs.cloudflare.com")
// }
