package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mooncake9527/orange-core/core"
	"github.com/mooncake9527/orange/common/consts"
)

var (
	routerNoCheckRole = make([]func(*gin.RouterGroup), 0)
	//routerCheckRole   = make([]func(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware), 0)
)

// InitRouter 路由初始化
func InitRouter() {
	r := core.GetGinEngine()
	noCheckRoleRouter(r)
}

// noCheckRoleRouter 无需认证的路由
func noCheckRoleRouter(r *gin.Engine) {
	// 可根据业务需求来设置接口版本
	v := r.Group(consts.ApiRoot + "/open")

	for _, f := range routerNoCheckRole {
		f(v)
	}
}
