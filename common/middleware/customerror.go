package middleware

import (
	stackUtil "github.com/mooncake9527/orange/common/utils/stack"
	"log/slog"
	"net/http"
	"runtime"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mooncake9527/orange-core/common/utils"
	"github.com/mooncake9527/orange-core/common/utils/ips"
	"github.com/mooncake9527/orange-core/core"
)

func CustomError(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			slog.Error("panic recover ", "error", err, "stack", stackUtil.GetStack())
			if c.IsAborted() {
				c.Status(200)
			}
			switch errStr := err.(type) {
			case string:
				p := strings.Split(errStr, "#")
				if len(p) == 3 && p[0] == "CustomError" {
					statusCode, e := strconv.Atoi(p[1])
					if e != nil {
						break
					}
					c.Status(statusCode)

					slog.Warn("request", "ip", ips.GetIP(c), "method", c.Request.Method, "path", c.Request.RequestURI,
						"query", c.Request.URL.RawQuery, "source", core.Cfg.Server.Name, "reqId", utils.GetReqId(c),
						"error", p[2])

					c.JSON(http.StatusOK, gin.H{
						"code": statusCode,
						"msg":  p[2],
					})
				} else {
					c.JSON(http.StatusOK, gin.H{
						"code": 500,
						"msg":  errStr,
					})
				}
			case runtime.Error:
				c.JSON(http.StatusOK, gin.H{
					"code": 500,
					"msg":  errStr.Error(),
				})
			default:
				panic(err)
			}
		}
	}()
	c.Next()
}
