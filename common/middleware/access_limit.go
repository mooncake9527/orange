package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/mooncake9527/orange-core/common/utils/ips"
	"github.com/mooncake9527/orange-core/core"
	"net/http"
	"time"
)

type Access struct {
	beginTime time.Time
	accessCnt int
}

var (
	accessMap = make(map[string]*Access, 0)
)

func AccessLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := ips.GetIP(c)
		v, ok := accessMap[ip]
		if !ok {
			accessMap[ip] = &Access{beginTime: time.Now(), accessCnt: 1}
		} else {
			curT := time.Now()
			if curT.Sub(v.beginTime) > core.Cfg.AccessLimit.Duration { //当前时间和开始时间的周期
				v.accessCnt = 1
				v.beginTime = curT
			} else if v.accessCnt > core.Cfg.AccessLimit.GetTotal() { //时间范围内数量超标
				v.accessCnt++
				if v.accessCnt/core.Cfg.AccessLimit.GetTotal() > 1 {
					v.beginTime = curT
				}
				Fail(c, 429, "too many requests")
				return
			} else {
				v.accessCnt++
			}
		}
		c.Next()
	}
}

func Fail(c *gin.Context, code int, msg string, data ...any) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}
