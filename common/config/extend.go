package config

import "github.com/mooncake9527/orange-core/config"

var Ext *Extend

func init() {
	Ext = new(Extend)
}

type Extend struct {
	Open      OpenMap  `mapstructure:"open" json:"open" yaml:"open"` // 这里配置对应配置文件的结构即可
	Ding      DingCfg  `mapstructure:"ding" json:"ding" yaml:"ding"`
	WechatMp  WechatMp `mapstructure:"wechat-mp" json:"wechat-mp" yaml:"wechat-mp"`
	WhileIps  string
	TestUser  string
	RdConfig  config.Config `mapstructure:"rd-config" json:"rd-config" yaml:"rd-config"`
	PayConfig PayConfig     `mapstructure:"pay-config" json:"pay-config" yaml:"pay-config"`
	AliOSS    AliOSS        `mapstructure:"ali-oss" json:"ali-oss" yaml:"ali-oss"`
}

type AliOSS struct {
	Endpoint        string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	AccessKeyId     string `mapstructure:"access-key-id" json:"access-key-id" yaml:"access-key-id"`
	AccessKeySecret string `mapstructure:"access-key-secret" json:"access-key-secret" yaml:"access-key-secret"`
	BucketName      string `mapstructure:"bucket-name" json:"bucket-name" yaml:"bucket-name"`
	BasePath        string `mapstructure:"base-path" json:"base-path" yaml:"base-path"`
	BucketUrl       string `mapstructure:"bucket-url" json:"bucket-url" yaml:"bucket-url"`
}

type PayConfig struct {
	AppId      string `mapstructure:"appid" json:"appid" yaml:"appid"`
	MerchantNo string `mapstructure:"merchantNo" json:"merchantNo" yaml:"merchantNo"`
	Currency   string `mapstructure:"currency" json:"currency" yaml:"currency"`
	NotifyUrl  string `mapstructure:"notify-url" json:"notify-url" yaml:"notify-url"`
	Url        string `mapstructure:"url" json:"url" yaml:"url"`
	PriKey     string `mapstructure:"priKey" json:"priKey" yaml:"priKey"`
	Prefix     string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
}

type DingCfg struct {
	AgentId   string `mapstructure:"agent-id" json:"agent-id" yaml:"agent-id"`
	AppKey    string `mapstructure:"app-key" json:"app-key" yaml:"app-key"`
	AppSecret string `mapstructure:"app-secret" json:"app-secret" yaml:"app-secret"`
	CropId    string `mapstructure:"crop-id" json:"crop-id" yaml:"crop-id"`
}

type WechatMp struct {
	AppId          string `mapstructure:"app-id" json:"app-id" yaml:"app-id"`
	AppSecret      string `mapstructure:"app-secret" json:"app-secret" yaml:"app-secret"`
	WxToken        string `mapstructure:"wx-token" json:"wx-token" yaml:"wx-token"`
	EncodingAESKey string `mapstructure:"encoding-aes-key" json:"encoding-aes-key" yaml:"encoding-aes-key"`
}

type OpenMap struct {
	GatewayUrl string `mapstructure:"gateway-url" json:"gateway-url" yaml:"gateway-url"`
	SsoUrl     string `mapstructure:"sso-url" json:"sso-url" yaml:"sso-url"`
}
