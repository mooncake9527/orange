package dto

import (
	"github.com/mooncake9527/orange-core/common/request"
	"github.com/mooncake9527/orange/modules/open/models"
)

type UserGetPageReq struct {
	// base.ReqPage `query:"-"`
	request.Pagination `query:"-"`
	SortOrder          string `json:"-" query:"type:order;column:id"`
}

func (x *UserGetPageReq) Valid() error {

	return nil
}

func (UserGetPageReq) TableName() string {
	return models.TBUser
}

// 用户表
type UserReq struct {
	Id       uint64 `json:"id" form:"id"`             //主键
	UserId   string `json:"userId" form:"userId"`     //用户id
	Username string `json:"username" form:"username"` //用户名
	Mobile   string `json:"mobile" form:"mobile"`     //手机号
	Nickname string `json:"nickname" form:"nickname"` //昵称
	Avatar   string `json:"avatar" form:"avatar"`     //头像
	Email    string `json:"email" form:"email"`       //邮箱
}

func (x *UserReq) Valid() error {

	return nil
}
