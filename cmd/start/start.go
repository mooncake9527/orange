package start

import (
	"fmt"
	"github.com/mooncake9527/orange-core/core"
	"github.com/mooncake9527/orange-core/core/ebus"
	"github.com/mooncake9527/orange-core/core/i18n"
	"github.com/mooncake9527/orange/cmd/com"
	"github.com/mooncake9527/orange/common/codes"
	"github.com/mooncake9527/orange/common/global"
	"github.com/mooncake9527/orange/common/middleware"
	stackUtil "github.com/mooncake9527/orange/common/utils/stack"
	"github.com/spf13/cobra"
	_ "github.com/spf13/viper/remote"
	"log/slog"
)

var (
	configYml string
	Cmd       = &cobra.Command{
		Use:     "start",
		Short:   "start application with config file",
		Example: "eg. orange start -c resources/config.yaml",
		Run: func(cmd *cobra.Command, args []string) {
			defer func() {
				if err := recover(); err != nil {
					slog.Error("panic err " + fmt.Sprintf("%v", err) + " \r\n stack:" + stackUtil.GetStack())
				}
			}()
			startApplication()
		},
	}
)

func init() {
	Cmd.PersistentFlags().StringVarP(&configYml, "config", "c", "resources/config.dev.yaml", "start server with provided configuration file")
}

func startApplication() {
	com.Pre(configYml)
	ebus.EventBus.Publish(global.TopicCoreInitFinish)
	i18n.Register(&codes.Code{
		EnableI18N: core.Cfg.Server.I18n,
		Lang:       core.Cfg.Server.Lang,
	})
	r := core.GetGinEngine()
	middleware.InitMiddleware(r, &core.Cfg)
	for _, f := range AppRouters {
		f()
	}
	go func() {
		<-core.Started
		startedInit()
	}()
	go func() {
		<-core.ToClose
		ebus.EventBus.Publish(global.TopicApplicationClose)
		toClose()
	}()
	core.Run()
	slog.Info("server exited")
}

func startedInit() {
	if core.Cfg.GrpcServer.Enable {
		grpcInit()
	}
	rdInit()
	slog.Debug("server started")
}

func toClose() {
	if core.Cfg.GrpcServer.Enable {
		closeGrpc()
	}
	rdRelease()
	slog.Debug("release resources")
}
