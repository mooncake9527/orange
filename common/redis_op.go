package common

import (
	"encoding/json"
	"fmt"
	jsonUtil "github.com/mooncake9527/orange/common/utils/json"
	"github.com/mooncake9527/orange/modules/open/models"
	"github.com/mooncake9527/x/xerrors/xerror"
	"github.com/spf13/cast"
	"time"

	"github.com/mooncake9527/orange-core/core"
)

const (
	RKOpen            = "r:open:"
	RKOpenApp         = RKOpen + "app:"
	RKOpenUser        = RKOpen + "user:"
	RKOpenAccount     = RKOpen + "account:"
	RKOpenApi         = RKOpen + "api:"
	RKOpenApiCount    = RKOpen + "apiCount:"
	RKOpenAccountGate = RKOpen + "account:gate:"

	RKOpenAccountGateOpen = 1
)

func OpenAccountGate(accountId uint64) error {
	id := cast.ToString(accountId)
	return core.Cache.Set(RKOpenAccountGate+id, RKOpenAccountGateOpen, -1)
}

func CloseAccountGate(accountId uint64) error {
	id := cast.ToString(accountId)
	return core.Cache.Del(RKOpenAccountGate + id)
}

func GetAccountGate(accountId uint64) (error, bool) {
	id := cast.ToString(accountId)
	j, err := core.Cache.Get(RKOpenAccountGate + id)
	if err != nil {
		return xerror.Wrap(err, "get AccountGate from cache error"), false
	}
	return nil, cast.ToBool(j)
}

// SetApp appSecret 如果不允许更改的话，expire时间可以写很长，或者直接放在memory
func SetApp(appId string, as models.App) {
	j, _ := json.Marshal(as)
	core.Cache.Set(RKOpenApp+appId, j, 7*24*time.Hour)
}

func DelApp(appId string) error {
	err := core.Cache.Del(RKOpenApp + appId)
	if err != nil {
		return xerror.Wrap(err, "delete AppSecret from cache error")
	}
	return nil
}

func GetApp(appId string, as *models.App) error {
	j, err := core.Cache.Get(RKOpenApp + appId)
	if err != nil {
		return xerror.Wrap(err, "get AppSecret from cache error")
	}
	if err = jsonUtil.Unmarshal([]byte(j), &as); err != nil {
		return err
	}
	return nil
}

func SetAccount(userId uint64, data models.Account) error {
	j, _ := json.Marshal(data)
	return core.Cache.Set(RKOpenAccount+cast.ToString(userId), j, 7*24*time.Hour)
}

func DelAccount(userId uint64) error {
	err := core.Cache.Del(RKOpenAccount + cast.ToString(userId))
	if err != nil {
		return xerror.Wrap(err, "delete Account from cache error")
	}
	return nil
}

func GetAccount(userId uint64, data *models.Account) error {
	j, err := core.Cache.Get(RKOpenAccount + cast.ToString(userId))
	if err != nil {
		return xerror.Wrap(err, "get Account from cache error")
	}
	if err = jsonUtil.Unmarshal([]byte(j), data); err != nil {
		return err
	}
	return nil
}

func SetApi(method, uri string, data models.Api) error {
	k := method + ":" + uri
	j, _ := json.Marshal(data)
	return core.Cache.Set(RKOpenApi+k, j, 7*24*time.Hour)
}

func DelApi(method, uri string) error {
	k := method + ":" + uri
	err := core.Cache.Del(RKOpenApi + k)
	if err != nil {
		return xerror.Wrap(err, "delete Api from cache error")
	}
	return nil
}

func GetApi(method, uri string, data *models.Api) error {
	k := method + ":" + uri
	j, err := core.Cache.Get(RKOpenApi + k)
	if err != nil {
		return xerror.Wrap(err, "get Api from cache error")
	}
	if err = jsonUtil.Unmarshal([]byte(j), data); err != nil {
		return err
	}
	return nil
}

func SetApiCount(appKey, method, uri string, data models.ApiCount) error {
	k := appKey + ":" + method + ":" + uri
	j, _ := json.Marshal(data)
	return core.Cache.Set(RKOpenApiCount+cast.ToString(k), j, 7*24*time.Hour)
}

func DelApiCount(appKey, method, uri string) error {
	k := appKey + ":" + method + ":" + uri
	err := core.Cache.Del(RKOpenApiCount + k)
	if err != nil {
		return xerror.Wrap(err, "delete ApiCount from cache error")
	}
	return nil
}

func GetApiCount(appKey, method, uri string, data *models.ApiCount) error {
	k := appKey + ":" + method + ":" + uri
	j, err := core.Cache.Get(RKOpenApiCount + k)
	if err != nil {
		return xerror.Wrap(err, "get ApiCount from cache error")
	}
	if err = jsonUtil.Unmarshal([]byte(j), data); err != nil {
		return err
	}
	return nil
}

func SetUser(userId string, user models.User) {
	j, _ := json.Marshal(user)
	core.Cache.Set(RKOpenUser+userId, j, 7*24*time.Hour)
}

func DelUser(userId string) error {
	err := core.Cache.Del(RKOpen + userId)
	if err != nil {
		return xerror.Wrap(err, "delete User from cache error")
	}
	return nil
}

func GetUser(userId string, user *models.User) error {
	j, err := core.Cache.Get(RKOpenUser + userId)
	if err != nil {
		return xerror.Wrap(err, "get User from cache error")
	}
	if err = jsonUtil.Unmarshal([]byte(j), &user); err != nil {
		return err
	}
	return nil
}

func GetMpAccessToken(appId string) string {
	j, err := core.Cache.Get("acctoken:" + appId)
	if err == nil {
		return j
	}
	return ""
}

func SetMpAccessToken(appId string, token string) {
	core.Cache.Set("acctoken:"+appId, token, 7000)
}

func GetMpOpenId(scene string) (string, error) {
	return core.Cache.Get("mp:login:" + scene)
}

func SetMpOpenId(scene, openId string) {
	core.Cache.Set("mp:login:"+scene, openId, 400)
}

func DelMpOpenId(scene string) error {
	err := core.Cache.Del("mp:login:" + scene)
	if err != nil {
		return err
	}
	return nil
}

func TeamMemberKey(teamId, userId int) string {
	return fmt.Sprintf("t:m:%d:%d", teamId, userId)
}
