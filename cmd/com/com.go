package com

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/jinzhu/copier"
	coreCfg "github.com/mooncake9527/orange-core/config"
	"github.com/mooncake9527/orange-core/core"
	"github.com/mooncake9527/orange/common/config"
	"github.com/spf13/viper"
	"time"
)

func Pre(configYml string) {
	if configYml == "" {
		panic("please specific config file")
	}
	v := viper.New()
	v.SetConfigFile(configYml)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("read config file error: %v", err))
	}

	var cfg coreCfg.AppCfg
	if err = v.Unmarshal(&cfg); err != nil {
		fmt.Println(err)
	}

	if cfg.Server.RemoteEnable {
		watchRemote(&cfg)
	} else {
		watchLocal(&cfg, v)
	}
	core.Init()
}

func watchRemote(cfg *coreCfg.AppCfg) {
	var err error
	iViper := viper.New()
	if cfg.Remote.SecretKeyring == "" {
		err = iViper.AddRemoteProvider(cfg.Remote.Provider, cfg.Remote.Endpoint, cfg.Remote.Path)
	} else {
		err = iViper.AddSecureRemoteProvider(cfg.Remote.Provider, cfg.Remote.Endpoint, cfg.Remote.Path, cfg.Remote.SecretKeyring)
	}
	if err != nil {
		panic(fmt.Sprintf("fetch remote config error: %v \n", err))
	}
	iViper.SetConfigType(cfg.Remote.GetConfigType())
	err = iViper.ReadRemoteConfig()
	if err != nil {
		panic(fmt.Sprintf("fetch remote config error: %v \n", err))
	}

	var remoteCfg coreCfg.AppCfg
	_ = iViper.Unmarshal(&remoteCfg)
	mergeCfg(cfg, &remoteCfg)
	extend := iViper.Sub("extend")
	if extend != nil {
		_ = extend.Unmarshal(config.Ext)
	}
	go func() {
		for {
			time.Sleep(time.Second * 5)
			err := iViper.WatchRemoteConfig()
			if err != nil {
				fmt.Println(err)
				continue
			}
			_ = iViper.Unmarshal(&remoteCfg)
			mergeCfg(cfg, &remoteCfg)
			extend := iViper.Sub("extend")
			if extend != nil {
				_ = extend.Unmarshal(config.Ext)
			}
		}
	}()
}

func watchLocal(cfg *coreCfg.AppCfg, v *viper.Viper) {
	var err error
	mergeCfg(cfg, nil)
	_ = v.Sub("extend").Unmarshal(config.Ext)
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.String())
		if err = v.Unmarshal(&cfg); err != nil {
			fmt.Println(err)
		}
		mergeCfg(cfg, nil)
		extend := v.Sub("extend")
		if extend != nil {
			_ = extend.Unmarshal(config.Ext)
		}
	})
}

func mergeCfg(local, remote *coreCfg.AppCfg) {
	if remote != nil {
		core.Cfg = *remote
		core.Cfg.Remote = local.Remote
		_ = copier.CopyWithOption(&core.Cfg.Server, &local.Server, copier.Option{IgnoreEmpty: true})
	} else {
		core.Cfg = *local
	}
}
