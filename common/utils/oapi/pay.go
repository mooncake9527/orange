package oapi

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/mooncake9527/orange-core/core/ebus"
	"github.com/mooncake9527/orange/common/config"
	"github.com/mooncake9527/orange/common/global"
	jsonUtil "github.com/mooncake9527/orange/common/utils/json"
	"github.com/mooncake9527/x/xerrors/xerror"
	"github.com/shopspring/decimal"
	"log/slog"
	"time"
)

// OrderReq 统一下单接口
type OrderReq struct {
	AppId        string `json:"appId" `        //商户Appid（终端号）
	MerchantNo   string `json:"merchantNo"`    //商户号
	AppTraceNo   string `json:"appTraceNo" `   //终端流水号(终端订单号)
	AppTraceTime int    `json:"appTraceTime" ` //终端交易时间
	TotalFee     string `json:"totalFee" `     //订单金额共36位，小数点18位
	Currency     string `json:"currency" `     //订单货币类型 RMB
	PayCurrency  string `json:"payCurrency" `  //支付货币类型 预留
	PayType      string `json:"payType" `      //支付方式
	OrderBody    string `json:"orderBody" `    //订单描述
	GoodsDetail  string `json:"goodsDetail" `  //订单包含的商品列表信息
	NotifyUrl    string `json:"notifyUrl"`     //回调地址
	RedirectUrl  string `json:"redirectUrl"`   //订单显示页
	Code         string `json:"code"`          //获取微信openId的code
	Payer        string `json:"payer"`         //支付人 wxjs支付openId
	DevType      string `json:"devType"`       //设备类型 h5支付
	Sign         string `json:"sign"`          //签名
	Locale       string `json:"locale"`        //语言
}

type OrderRsp struct {
	ReqId string
	Code  int
	Msg   string
	Data  interface{}
}

// OrderPreResp 预付订单返回
type OrderPreResp struct {
	OutTradeNo  string `json:"outTradeNo" ` //系统平台唯一订单号
	AppTraceNo  string `json:"appTraceNo" ` //终端流水号(终端订单号)
	OrderStatus int    `json:"orderStatus"` //订单状态 例如1待付款，2已付款，3已取消等 4异常
	Currency    string `json:"currency" `   //货币类型 RMB
	PayType     string `json:"payType"`     //支付方式
	PayTrace    string `json:"payTrace" `   //实际支付金额（小数点18位）
	PayUrl      string `json:"payUrl" `     //支付链接
	Ext         any    `json:"ext"`         //扩张
}

// OrderResp 订单回调信息
type OrderResp struct {
	MerchantNo string `json:"merchantNo" ` //商户号
	AppId      string `json:"appId" `      //商户Appid（终端号）
	PayType    string `json:"payType"`     //支付方式
	Body       string `json:"body"`        //订单信息 OrderInfoResp
}

// OrderInfoResp 回调订单信息
type OrderInfoResp struct {
	OutTradeNo  string      `json:"outTradeNo" ` //支付系统平台唯一订单号
	AppTraceNo  string      `json:"appTraceNo" ` //终端流水号(终端订单号) 本系统订单号
	OrderStatus OrderStatus `json:"orderStatus"` //订单状态 例如1待付款，2已付款，3已取消等 4异常
	PayOrderNo  string      `json:"payOrderNo" ` //支付平台订单号 第三方
	MerchantNo  string      `json:"merchantNo" ` //商户号
	AppId       string      `json:"appId" `      //商户Appid（终端号）
	Currency    string      `json:"currency" `   //货币类型 RMB
	PayType     string      `json:"payType"`     //支付方式
	PayTrace    string      `json:"payTrace" `   //实际支付金额（小数点18位）
	PayTime     int64       `json:"payTime" `    //支付平台 支付时间
}

type OrderStatus int

const (
	OrderStatusWaitPay    OrderStatus = 1
	OrderStatusPaySuccess OrderStatus = 2
	OrderStatusCancel     OrderStatus = 3
	OrderStatusException  OrderStatus = 4
)

// GetOrderReq 查询订单请求
type GetOrderReq struct {
	AppId      string `json:"appId" `      //商户Appid（终端号）
	MerchantNo string `json:"merchantNo"`  //商户号
	AppTraceNo string `json:"appTraceNo" ` //终端流水号(终端订单号)
	OutTradeNo string `json:"outTradeNo" ` //系统平台唯一订单号
	ReqTime    int    `json:"reqTime" `    //请求时间
	Sign       string `json:"sign"`        //签名 Rsa appTraceNo=%s&merchantNo=%s&outTradeNo=%s&ReqTime=%d
}

func CreatePay(orderNo, payType, subject, user, redirectUrl, payer string, totalFee decimal.Decimal) (orderPreResp OrderPreResp, err error) {
	goodsDetail := ""
	finalAmount := totalFee.String()
	t := time.Now()
	reqOrder := OrderReq{
		AppId:        config.Ext.PayConfig.AppId,
		MerchantNo:   config.Ext.PayConfig.MerchantNo,
		AppTraceNo:   orderNo,
		Locale:       "",
		AppTraceTime: int(t.Unix()),
		TotalFee:     finalAmount,
		Currency:     config.Ext.PayConfig.Currency,
		PayType:      payType, //payTyp := "wx","zfb","abc","t"(聚合), "wxjs", "wxh5", "wxApp", "aliApp", "alimini", "wxmini", "airw", "paypal", "strh5", "ethusd", "trxusd"
		OrderBody:    subject,
		GoodsDetail:  goodsDetail,
		NotifyUrl:    config.Ext.PayConfig.NotifyUrl,
		RedirectUrl:  redirectUrl,
		Payer:        payer,
	}
	str, err := GenSign(reqOrder)
	if err != nil {
		return orderPreResp, xerror.Wrap(err, "请求支付加密失败CreatePay")
	}
	reqOrder.Sign = EncodeToString(str)
	reqData, _ := json.Marshal(reqOrder)
	slog.Info("CreatePay", "req", string(reqData))
	response, err := New(config.Ext.PayConfig.Url).AddHeader("Content-Type", "application/json").Post("/v2/payment/createOrder", reqData)
	ebus.EventBus.Publish(global.TopicTrafficLog, reqData, response, orderNo)
	if err != nil {
		return orderPreResp, xerror.Wrap(err, "请求支付失败CreatePay")
	}
	slog.Info("CreatePay", "rsp", string(response))
	var orderRsp OrderRsp
	_ = jsonUtil.Unmarshal(response, &orderRsp)
	if orderRsp.Code != 200 {
		return orderPreResp, xerror.New("请求支付返回失败CreatePay")
	} else {
		if err = mapstructure.Decode(orderRsp.Data, &orderPreResp); err != nil {
			return orderPreResp, xerror.Wrap(err, "请求支付解析失败CreatePay")
		}
	}
	return orderPreResp, nil
}

func GetOrder(ctx context.Context, orderNo string) (orderFindResp OrderInfoResp, err error) {
	t := time.Now()
	order := GetOrderReq{
		AppId:      config.Ext.PayConfig.AppId,
		MerchantNo: config.Ext.PayConfig.MerchantNo,
		AppTraceNo: orderNo,
		ReqTime:    int(t.Unix()),
	}
	s, err := GenGetSign(order)
	if err != nil {
		return
	}
	order.Sign = EncodeToString(s)
	data, _ := jsonUtil.Marshal(order)
	rsp, err := New(config.Ext.PayConfig.Url).AddHeader("Content-Type", "application/json").Post("/v2/payment/get", data)
	if err != nil {
		return orderFindResp, xerror.Wrap(err, "请求支付查询失败GetOrder")
	}
	var orderRsp OrderRsp
	_ = jsonUtil.Unmarshal(rsp, &orderRsp)
	if orderRsp.Code != 200 {
		return orderFindResp, xerror.New("请求支付返回失败GetOrder")
	} else {
		if err = mapstructure.Decode(orderRsp.Data, &orderFindResp); err != nil {
			return orderFindResp, xerror.Wrap(err, "请求支付查询解析失败GetOrder")
		}
	}
	return orderFindResp, nil
}

func GenSign(req OrderReq) ([]byte, error) {
	str := fmt.Sprintf("appTraceNo=%s&appTraceTime=%d&currency=%s&merchantNo=%s&payCurrency=%s&payType=%s&totalFee=%s", req.AppTraceNo, req.AppTraceTime, req.Currency, req.MerchantNo, req.PayCurrency, req.PayType, req.TotalFee)
	sign, err := RSA_Sign(config.Ext.PayConfig.PriKey, []byte(str))
	if err != nil {
		return nil, xerror.Wrap(err, "pay sign err")
	}
	return sign, nil
}

func GenGetSign(req GetOrderReq) ([]byte, error) {
	str := fmt.Sprintf("appTraceNo=%s&merchantNo=%s&outTradeNo=%s&ReqTime=%d", req.AppTraceNo, req.MerchantNo, req.OutTradeNo, req.ReqTime)
	sign, err := RSA_Sign(config.Ext.PayConfig.PriKey, []byte(str))
	if err != nil {
		fmt.Println("sign err")
		fmt.Println(err)
	}
	return sign, nil
}
