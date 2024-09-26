package middleware

import (
	"errors"
	"fmt"
	"github.com/mooncake9527/orange-core/common/utils"
	gLocal "github.com/mooncake9527/orange-core/common/xlog/g_local"
	"github.com/mooncake9527/orange-core/core"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var (
	Secret              = core.Cfg.JWT.SignKey                              //密码自行设定
	TokenExpireDuration = time.Duration(core.Cfg.JWT.Expires) * time.Minute //设置过期时间

	TokenLookup   = "header: Authorization, query: token, cookie: jwt"
	TokenHeadName = "Bearer"

	// ErrEmptyAuthHeader can be thrown if authing with a HTTP header, the Decrypt header needs to be set
	ErrEmptyAuthHeader = errors.New("auth header is empty")

	// ErrInvalidAuthHeader indicates auth header is invalid, could for example have the wrong Realm name
	ErrInvalidAuthHeader = errors.New("auth header is invalid")

	// ErrEmptyQueryToken can be thrown if authing with URL Query, the query token variable is empty
	ErrEmptyQueryToken = errors.New("query token is empty")

	// ErrEmptyCookieToken can be thrown if authing with a cookie, the token cokie is empty
	ErrEmptyCookieToken = errors.New("cookie token is empty")

	// ErrEmptyParamToken can be thrown if authing with parameter in path, the parameter in path is empty
	ErrEmptyParamToken = errors.New("parameter token is empty")
)

func JWTAuthMiddleware() func(c *gin.Context) {
	if Secret == "" {
		Secret = core.Cfg.JWT.SignKey                                           //密码自行设定
		TokenExpireDuration = time.Duration(core.Cfg.JWT.Expires) * time.Minute //设置过期时间
	}
	return func(c *gin.Context) {
		var token string
		var err error

		methods := strings.Split(TokenLookup, ",")
		for _, method := range methods {
			if len(token) > 0 {
				break
			}
			parts := strings.Split(strings.TrimSpace(method), ":")
			k := strings.TrimSpace(parts[0])
			v := strings.TrimSpace(parts[1])
			switch k {
			case "header":
				token, err = jwtFromHeader(c, v)
			case "query":
				token, err = jwtFromQuery(c, v)
			case "cookie":
				token, err = jwtFromCookie(c, v)
			case "param":
				token, err = jwtFromParam(c, v)
			}
		}

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 1001,
				"msg":  "无有效token",
			})
			c.Abort()
			return
		}

		mc, err := ParseToken(token, Secret)

		if err != nil {
			slog.Info(fmt.Sprintf("无效的Token userId:%s token:%s exp:%s", mc.Id, token, time.Unix(mc.ExpiresAt, 0).Format("2006-01-02 15:04:05")))
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 1001,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}

		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("userId", mc.Id)
		gLocal.SetCUId(mc.Id, "")
		defer gLocal.SetCUId("", "")
		if mc.Issuer != "" {
			c.Set("companyId", mc.Issuer)
		}

		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}

func ParseToken(tokenString string, secret string) (jwt.StandardClaims, error) {
	sc := jwt.StandardClaims{}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return sc, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid { // 校验token
		if value, ok := claims["exp"]; ok {
			sc.ExpiresAt = int64(utils.GetInterfaceToInt(value))
		}
		if value, ok := claims["jti"]; ok {
			sc.Id = value.(string)
		}
		if value, ok := claims["iss"]; ok {
			sc.Issuer = value.(string)
		}
		return sc, nil
	}

	return sc, errors.New("invalid token")
}

func GenToken(userId, companyId string, timeout time.Time, secret string) (string, error) {
	c := jwt.StandardClaims{
		Id:        userId,
		ExpiresAt: timeout.Unix(),
	}
	if companyId != "" {
		c.Issuer = companyId
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenString, err := token.SignedString([]byte(secret))
	return tokenString, err
}

func jwtFromHeader(c *gin.Context, key string) (string, error) {
	authHeader := c.Request.Header.Get(key)
	if authHeader == "" {
		return "", ErrEmptyAuthHeader
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == TokenHeadName) {
		return "", ErrInvalidAuthHeader
	}

	return parts[1], nil
}

func jwtFromQuery(c *gin.Context, key string) (string, error) {
	token := c.Query(key)

	if token == "" {
		return "", ErrEmptyQueryToken
	}

	return token, nil
}

func jwtFromCookie(c *gin.Context, key string) (string, error) {
	cookie, _ := c.Cookie(key)

	if cookie == "" {
		return "", ErrEmptyCookieToken
	}

	return cookie, nil
}

func jwtFromParam(c *gin.Context, key string) (string, error) {
	token := c.Param(key)

	if token == "" {
		return "", ErrEmptyParamToken
	}

	return token, nil
}
