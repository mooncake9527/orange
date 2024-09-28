package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mooncake9527/orange/modules/open/apis"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerUserRouter)
}

// 默认需登录认证的路由
func registerUserRouter(v1 *gin.RouterGroup) {
	r := v1.Group("user") //.Use(middleware.JwtHandler()).Use(middleware.PermHandler())
	{
		r.POST("/get", apis.ApiUser.Get)
		r.POST("/create", apis.ApiUser.Create)
		r.POST("/update", apis.ApiUser.Update)
		r.POST("/page", apis.ApiUser.QueryPage)
		r.POST("/del", apis.ApiUser.Del)
	}
}
