package models

import (
	"time"
)

// User 用户表
type User struct {
	Id        uint64     `json:"id"  form:"id"  gorm:"type:bigint unsigned;primaryKey;autoIncrement;comment:主键"`         //主键
	UserId    string     `json:"userId" form:"userId" gorm:"type:varchar(64);comment:用户id"`                              //用户id
	Username  string     `json:"username" form:"username" gorm:"type:varchar(50);comment:用户名"`                           //用户名
	Mobile    string     `json:"mobile" form:"mobile" gorm:"type:varchar(18);comment:手机号"`                               //手机号
	Nickname  string     `json:"nickname" form:"nickname" gorm:"type:varchar(64);comment:昵称"`                            //昵称
	Avatar    string     `json:"avatar" form:"avatar" gorm:"type:varchar(255);comment:头像"`                               //头像
	Email     string     `json:"email" form:"email" gorm:"type:varchar(128);comment:邮箱"`                                 //邮箱
	CreatedAt time.Time  `json:"createdAt" form:"createdAt" gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:创建时间"` //创建时间
	UpdatedAt *time.Time `json:"updatedAt" form:"updatedAt" gorm:"type:datetime;comment:更新时间"`                           //更新时间
}

const TBUser = "user"

func (User) TableName() string {
	return TBUser
}
