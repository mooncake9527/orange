package start

import (
	"fmt"
	"github.com/mooncake9527/orange-core/rd"
	"github.com/mooncake9527/orange/common/config"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mooncake9527/orange-core/common/utils/ips"
	"github.com/mooncake9527/orange-core/core"
)

var (
	routerNoCheckRole = make([]func(*gin.RouterGroup), 0)
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerHealthRouter)
	AppRouters = append(AppRouters, InitRouter)
}

func InitRouter() {
	r := core.GetGinEngine()
	noCheckRoleRouter(r)
}

func noCheckRoleRouter(r *gin.Engine) {
	v := r.Group("")

	for _, f := range routerNoCheckRole {
		f(v)
	}
}

func registerHealthRouter(v1 *gin.RouterGroup) {
	r := v1.Group("")
	{
		r.GET("/api/health", func(ctx *gin.Context) {
			ctx.AbortWithStatus(http.StatusOK)
		})
	}
}

var rdclient rd.RDClient

func rdInit() {
	//注册中心
	if config.Ext.RdConfig.Enable {
		rdcfg := config.Ext.RdConfig
		for _, v := range rdcfg.Registers {
			if v.Protocol != "grpc" && v.Protocol != "http" {
				slog.Error("rd register error", "protocol", v.Protocol)
				continue
			}
			if v.Protocol == "grpc" && !core.Cfg.GrpcServer.Enable {
				slog.Error("rd register error", "protocol", v.Protocol, "GrpcServer enable", false)
				continue
			}
			if v.Name == "" {
				if v.Protocol == "http" {
					v.Name = core.Cfg.Server.Name
				} else {
					if core.Cfg.GrpcServer.Name != "" {
						v.Name = core.Cfg.GrpcServer.Name
					} else {
						v.Name = core.Cfg.Server.Name + "_grpc"
					}
				}
			}
			if v.Addr == "" {
				if v.Protocol == "http" {
					if core.Cfg.Server.GetHost() != "0.0.0.0" {
						v.Addr = core.Cfg.Server.GetHost()
					} else {
						v.Addr = ips.GetLocalHost()
					}
					v.Port = core.Cfg.Server.GetPort()
					v.HealthCheck = fmt.Sprintf("http://%s:%d/api/health", v.Addr, core.Cfg.Server.GetPort())
				} else {
					if core.Cfg.GrpcServer.GetHost() != "0.0.0.0" {
						v.Addr = core.Cfg.GrpcServer.GetHost()
					} else {
						v.Addr = ips.GetLocalHost()
					}
					v.Port = core.Cfg.GrpcServer.GetPort()
					v.HealthCheck = fmt.Sprintf("%s:%d/Health", v.Addr, core.Cfg.GrpcServer.GetPort())
				}
			}
			if len(v.Tags) == 0 {
				v.Tags = []string{core.Cfg.Server.Mode}
			}
			if v.Id == "" {
				v.Id = fmt.Sprintf("%s:%d", v.Addr, v.Port)
			}
		}

		slog.Debug("注册中心连接", "rdcfg", rdcfg)
		var err error
		rdclient, err = rd.NewRDClient(&rdcfg)
		if err != nil {
			slog.Error("注册中心连接失败", err)
		}
	}
}

func rdRelease() {
	if rdclient != nil {
		rdclient.Deregister()
	}
}
