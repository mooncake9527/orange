package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/mooncake9527/orange/common/consts"
	ctxUtil "github.com/mooncake9527/orange/common/utils/ctx"
)

func ValidUserId() func(*gin.Context) {
	return func(c *gin.Context) {
		userId := ctxUtil.GetUserId(c)
		if userId == "" {
			Fail(c, consts.ErrUnLogin.Code(), consts.ErrUnLogin.Error())
			return
		}
		c.Next()
	}
}
