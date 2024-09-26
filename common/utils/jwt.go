package utils

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mooncake9527/orange-core/common/utils"
)

// GetUserId 获取用户的uuid
func GetUserId(c *gin.Context) string {
	userId, exists := c.Get("userId")
	if !exists {
		return ""
	}
	if userIdStr, ok := userId.(string); ok {
		return userIdStr
	}
	return ""
}

func GetSub(c *gin.Context) int {
	sub, exists := c.Get("sub")
	if !exists {
		return 0
	}
	if userIdStr, ok := sub.(int); ok {
		return userIdStr
	}
	return 0
}

// GetCompanyId 获取企业company的uuid
func GetCompanyId(c *gin.Context) string {
	companyId, exists := c.Get("companyId")
	if !exists {
		return ""
	}
	if companyIdStr, ok := companyId.(string); ok {
		return companyIdStr
	}
	return ""
}

// AdminCustomClaims 自定义格式内容
type CustomClaims struct {
	UserId               int    `json:"uid,omitempty"`
	RoleId               int    `json:"rid,omitempty"`
	Phone                string `json:"mob,omitempty"`
	Nickname             string `json:"nick,omitempty"`
	JwtData              map[string]any
	jwt.RegisteredClaims // 内嵌标准的声明
}

func (c *CustomClaims) AddData(key string, val any) *CustomClaims {
	if c.JwtData == nil {
		c.JwtData = make(map[string]any, 0)
	}
	c.JwtData[key] = val
	return c
}

func (c *CustomClaims) GetInt(key string) int {
	if val, ok := c.JwtData[key]; ok {
		return utils.GetInterfaceToInt(val)
	}
	return 0
}

func (c *CustomClaims) GetString(key string) string {
	if val, ok := c.JwtData[key]; ok {
		return fmt.Sprintf("%s", val)
	}
	return ""
}

func (c *CustomClaims) ExpiresAt(expiresAt time.Time) *CustomClaims {
	c.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(expiresAt)
	return c
}

// NewAdminCustomClaims 初始化AdminCustomClaims
func NewClaims(userId int, expiresAt time.Time, issuer, subject string) CustomClaims {
	//now := time.Now()
	return CustomClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt), // 定义过期时间
			Issuer:    issuer,                        // 签发人
			//IssuedAt:  jwt.NewNumericDate(now),       // 签发时间
			Subject: subject, // 签发主体
			//NotBefore: jwt.NewNumericDate(now),       // 生效时间
		},
	}
}
