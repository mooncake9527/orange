package middleware

import (
	"bytes"
	stringUtil "github.com/mooncake9527/orange/common/utils/string"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mooncake9527/orange-core/common/utils"
	"github.com/mooncake9527/orange-core/common/utils/ips"
	"github.com/mooncake9527/orange-core/core"
)

// LogLayout 日志layout
// type LogLayout struct {
// 	//Metadata  map[string]interface{} // 存储自定义原数据
// 	Method    string //方法
// 	Path      string // 访问路径
// 	Query     string // 携带query
// 	Body      string // 携带body数据
// 	IP        string // ip地址
// 	UserAgent string // 代理
// 	Error     string // 错误
// 	Cost      string // 花费时间
// 	Source    string // 来源
// }

// LoggerToFile 日志记录到文件

func LoggerToFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		var body string
		method := strings.ToUpper(c.Request.Method)
		switch method {
		case http.MethodPost, http.MethodPut, http.MethodGet, http.MethodDelete:
			reqBodyBytes, _ := io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(reqBodyBytes))
			body = string(reqBodyBytes)
		}

		slog.Info("request", "method", c.Request.Method, "path", c.Request.RequestURI,
			"reqBody", body, "reqId", utils.GetReqId(c))

		crw := NewCustomResponseWriter(c.Writer)
		c.Writer = crw

		c.Next()

		writeLog(startTime, body, crw.body.String(), c)
	}
}

type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func NewCustomResponseWriter(responseWriter gin.ResponseWriter) *CustomResponseWriter {
	return &CustomResponseWriter{body: bytes.NewBufferString(""), ResponseWriter: responseWriter}
}

func (w CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

const limit = 128

func writeLog(startTime time.Time, body, respBody string, c *gin.Context) {
	// 结束时间
	if c.Request.Method == http.MethodOptions {
		return
	}
	cost := time.Since(startTime)

	if cost.Milliseconds() < 200 {
		slog.Info("request", "ip", ips.GetIP(c), "method", c.Request.Method, "path", c.Request.RequestURI,
			"cost", cost, "userAgent", c.Request.UserAgent(), "query", c.Request.URL.RawQuery,
			"reqBody", stringUtil.Limit(body, limit), "respBody", stringUtil.Limit(respBody, limit), "source", core.Cfg.Server.Name, "reqId", utils.GetReqId(c))
		//,"error", strings.TrimRight(c.Errors.ByType(gin.ErrorTypePrivate).String(), "\n")))
	} else {
		slog.Warn("request", "ip", ips.GetIP(c), "method", c.Request.Method, "path", c.Request.RequestURI,
			"cost", cost, "userAgent", c.Request.UserAgent(), "query", c.Request.URL.RawQuery,
			"reqBody", stringUtil.Limit(body, limit), "respBody", stringUtil.Limit(respBody, limit), "source", core.Cfg.Server.Name, "reqId", utils.GetReqId(c))
		//,"error", strings.TrimRight(c.Errors.ByType(gin.ErrorTypePrivate).String(), "\n")))
	}
}
