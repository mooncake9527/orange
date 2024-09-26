package oapi

import (
	"fmt"
	"github.com/mooncake9527/orange-core/common/xlog/xlog"
	"github.com/mooncake9527/orange/common/consts"
	hashUtil "github.com/mooncake9527/orange/common/utils/hash"
	"github.com/mooncake9527/orange/common/utils/http"
	jsonUtil "github.com/mooncake9527/orange/common/utils/json"
	stringUtil "github.com/mooncake9527/orange/common/utils/string"
	uuidUtil "github.com/mooncake9527/orange/common/utils/uuid"
	"github.com/mooncake9527/orange/modules/open/service/dto"
	"github.com/mooncake9527/x/xerrors/xerror"
	"log/slog"
	"net/http"
	"sort"
	"strings"
	"time"
)

var oApiMap = map[string]*OApi{}

type OApi struct {
	client *httpUtil.HTTPClient
}

func New(baseUrl string) *OApi {
	if oa, ok := oApiMap[baseUrl]; ok {
		return oa
	}
	oa := &OApi{
		client: httpUtil.New(baseUrl),
	}
	oApiMap[baseUrl] = oa
	return oa
}

func NewFingerOApi(baseUrl string) *OApi {
	if oa, ok := oApiMap[baseUrl]; ok {
		return oa
	}
	oa := &OApi{
		client: httpUtil.New(baseUrl),
	}
	oApiMap[baseUrl] = oa
	return oa
}

type CommonRes struct {
	ReqId string `json:"reqId,omitempty"` //`json:"请求id"`
	Code  int    `json:"code"`            //返回码
	Msg   string `json:"msg,omitempty"`   //消息
	Data  any    `json:"data,omitempty"`  //数据
}

type CreateFingerReq struct {
	System    string `json:"system" form:"system"`       //
	UAversion string `json:"uAversion" form:"uAversion"` //
	PublicIp  string `json:"publicIp" form:"publicIp"`   //
	IpChannel string `json:"ipChannel" form:"ipChannel"` //
	Kernel    string `json:"kernel" form:"kernel"`       //
}

type CreateFingerResp struct {
	ReqId string `json:"reqId,omitempty"` //`json:"请求id"`
	Code  int    `json:"code"`            //返回码
	Msg   string `json:"msg,omitempty"`   //消息
	Data  struct {
		DefaultConfig any `json:"defaultConfig"`
		FingerInfo    any `json:"fingerInfo"`
	} `json:"data,omitempty"` //数据
}

func (x *OApi) CreateFinger(req, resp any) error {
	bytes, _ := jsonUtil.Marshal(req)
	params := make(map[string]string)
	_ = jsonUtil.Unmarshal(bytes, &params)
	sign := fingerSign(params)
	params["sign"] = sign
	reqBytes, _ := jsonUtil.Marshal(params)
	response, err := x.client.Send(fingerCreateUri, http.MethodPost, reqBytes)
	if err != nil {
		slog.Error("http send fail", "error", err)
		return err
	}
	xlog.Info("http send success", "resp", stringUtil.Limit(string(response), 256))
	err = jsonUtil.Unmarshal(response, resp)
	if err != nil {
		res := CreateFingerResp{
			ReqId: uuidUtil.Gen(),
			Code:  500,
			Msg:   string(response),
			Data: struct {
				DefaultConfig any `json:"defaultConfig"`
				FingerInfo    any `json:"fingerInfo"`
			}{},
		}
		bs, _ := jsonUtil.Marshal(res)
		_ = jsonUtil.Unmarshal(bs, resp)
		return consts.Err500
	}
	return nil
}

func (e *OApi) Header(m map[string]string) *OApi {
	for k, v := range m {
		e.client.Headers[k] = v
	}
	return e
}

func (e *OApi) AddHeader(k, v string) *OApi {
	e.client.Headers[k] = v
	return e
}

const (
	fingerCreateUri = "/v2/userapi/open/generateFinger"
	fingerSecretKey = "jtSElesnJyJrTlqqFhrl"
)

func fingerSign(params map[string]string) string {
	keys := make([]string, 0, len(params))
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	signStr := ""
	for _, v := range keys {
		signStr += fmt.Sprintf("%v=%v&", v, params[v])
	}
	signStr += fmt.Sprintf("secretKey=%v", fingerSecretKey)
	return hashUtil.MD5([]byte(signStr))
}

const userUri = "/v2/sso/getUserinfo"
const userUriUpdate = "/v2/sso/getUserinfoUpdate"
const appVersion = "/v2/team/openVersion/get"

type Res struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Userinfo struct {
	UserId     string       `json:"userId" comment:"用户id"`
	Username   string       `json:"username" comment:"用户名"`
	Mobile     string       `json:"mobile" comment:"手机号"`
	Email      string       `json:"email" comment:"邮箱"`
	FirstName  string       `json:"firstName" comment:"名"`
	LastName   string       `json:"lastName" comment:"姓"`
	Nickname   string       `json:"nickname" comment:"昵称"`
	Avatar     string       `json:"avatar" comment:"头像"`
	Bio        string       `json:"bio" comment:"签名"`
	Gender     string       `json:"gender" comment:"性别 0 女 1 男 2 未知"`
	Birthday   time.Time    `json:"birthday" comment:"生日"`
	Inviter    string       `json:"inviter" gorm:"type:varchar(32);default:(-);comment:邀请人"` //邀请人
	InviteType int          `json:"inviteType" gorm:"type:tinyint;comment:邀请类型"`             //邀请类型
	OpenType   OpenTypeEnum `json:"openType" gorm:"type:tinyint;comment:开放类型"`               //开放类型
	CreatedAt  time.Time    `json:"createdAt" gorm:"comment:创建时间"`                           //创建时间
	Source     string       `json:"source" gorm:"type:varchar(32);default:(-);comment:来源"`   //来源
}

type OpenTypeEnum int

const (
	OpenPlatform OpenTypeEnum = 1
)

func (u Userinfo) GetName() string {
	if u.LastName != "" && u.FirstName != "" {
		return u.LastName + u.FirstName
	} else if u.FirstName != "" {
		return u.FirstName
	} else if u.LastName != "" {
		return u.LastName
	} else if u.Nickname != "" {
		return u.Nickname
	} else if u.Mobile != "" {
		return u.Mobile
	} else if u.Email != "" {
		arr := strings.Split(u.Email, "@")
		return arr[0]
	}
	return u.Username
}

func (e *OApi) GetUserInfo(userId string, user *Userinfo) error {
	m := map[string]string{
		"id": userId,
	}
	data, err := jsonUtil.Marshal(m)
	if err != nil {
		return err
	}
	response, err := e.client.Post(userUri, data)
	if err != nil {
		xlog.Error("GetUserInfo fail", "err", err.Error(), "req", string(data), "resp", string(response))
		return xerror.New(err.Error())
	}
	xlog.Info("GetUserInfo success", "req", string(data), "resp", string(response))
	fmt.Println(string(response))
	var res Res
	if err := jsonUtil.Unmarshal(response, &res); err != nil {
		return xerror.New(err.Error())
	}
	if res.Code == 200 {
		d, err := jsonUtil.Marshal(res.Data)
		if err != nil {
			return xerror.New(err.Error())
		}
		return jsonUtil.Unmarshal(d, user)
	} else {
		return xerror.New(res.Msg)
	}
}

func (e *OApi) GetUserInfoUpdate(userId string, user *Userinfo) error {
	m := map[string]string{
		"id": userId,
	}
	data, err := jsonUtil.Marshal(m)
	if err != nil {
		return err
	}
	response, err := e.client.Post(userUriUpdate, data)
	if err != nil {
		xlog.Error("GetUserInfoUpdate fail", "err", err.Error(), "req", string(data), "resp", string(response))
		return xerror.New(err.Error())
	}
	xlog.Info("GetUserInfoUpdate success", "req", string(data), "resp", string(response))
	fmt.Println(string(response))
	var res Res
	if err := jsonUtil.Unmarshal(response, &res); err != nil {
		return xerror.New(err.Error())
	}
	if res.Code == 200 {
		d, err := jsonUtil.Marshal(res.Data)
		if err != nil {
			return xerror.New(err.Error())
		}
		return jsonUtil.Unmarshal(d, user)
	} else {
		return xerror.New(res.Msg)
	}
}

type AppVersionRes struct {
	Code int                   `json:"code"`
	Msg  string                `json:"msg"`
	Data dto.GetAppVersionResp `json:"data"`
}

func (e *OApi) AppVersion() (*dto.GetAppVersionResp, error) {
	params := map[string]string{}
	data, err := jsonUtil.Marshal(params)
	if err != nil {
		return nil, err
	}
	response, err := e.client.Post(appVersion, data)
	if err != nil {
		xlog.Error("AppVersion fail", "err", err.Error(), "req", string(data), "resp", string(response))
		return nil, xerror.New(err.Error())
	}
	xlog.Info("AppVersion success", "req", string(data), "resp", string(response))
	var res AppVersionRes
	if err := jsonUtil.Unmarshal(response, &res); err != nil {
		return nil, xerror.New(err.Error())
	}
	if res.Code == 200 {
		return &res.Data, nil
	} else {
		return nil, xerror.New(res.Msg)
	}
}

func (e *OApi) Post(url string, data []byte) (resp []byte, err error) {
	resp, err = e.client.Post(url, data)
	if err != nil {
		xlog.Error("Post fail", "err", err.Error(), "req", string(data), "resp", string(resp))
		return nil, xerror.New(err.Error())
	}
	return resp, nil
}
func (e *OApi) Get(url string) (resp []byte, err error) {
	resp, err = e.client.Get(url)
	if err != nil {
		xlog.Error("get fail", "err", err.Error(), "resp", string(resp))
		return nil, xerror.New(err.Error())
	}
	return resp, nil
}
