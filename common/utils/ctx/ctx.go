package ctxUtil

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mooncake9527/orange-core/common/consts"
)

const (
	CtxKeyAppKey    = "appKey"
	CtxKeyUserId    = "userId"
	CtxKeySub       = "sub"
	CtxKeyCompanyId = "companyId"
	ctxKeyUser      = "open:user:"
)

// GetUserId 获取用户的uuid
func GetUserId(c *gin.Context) string {
	return getFromCtx[string](c, CtxKeyUserId)
}

func GetAppKey(c *gin.Context) string {
	return getFromCtx[string](c, CtxKeyAppKey)
}

func SetAppKey(c *gin.Context, appKey string) {
	c.Set(CtxKeyAppKey, appKey)
}

func GetReqId(c *gin.Context) string {
	reqId := c.GetString(consts.ReqId)
	if reqId == "" {
		reqId = uuid.NewString()
		c.Set(consts.ReqId, reqId)
	}
	return reqId
}

func GetSub(c *gin.Context) int {
	return getFromCtx[int](c, CtxKeySub)
}

// GetCompanyId 获取企业company的uuid
func GetCompanyId(c *gin.Context) string {
	return getFromCtx[string](c, CtxKeyCompanyId)
}

func getFromCtx[T any](c *gin.Context, key string) T {
	var dft T // default value
	val, exists := c.Get(key)
	if !exists {
		return dft
	}
	if v, ok := val.(T); ok {
		return v
	}
	return dft
}

func tryGetFromCtx[T any](c *gin.Context, keys []string) T {
	var ret T
	for _, key := range keys {
		val, exists := c.Get(key)
		if exists {
			return val.(T)
		}
	}
	return ret
}
