package httpUtil

import (
	"bytes"
	"github.com/mooncake9527/orange-core/common/xlog/xlog"
	"github.com/mooncake9527/x/xerrors/xerror"
	"io"
	"net/http"
)

func New(addr string, opts ...HTTPOption) *HTTPClient {
	o := defaultHTTPOptions()
	o.apply(opts...)
	c := &HTTPClient{
		baseURL: addr,
	}
	c.Headers = Headers{}
	c.Headers[HeaderContentType] = o.contentType
	c.sign = o.sign
	return c
}

type HTTPClient struct {
	baseURL string
	Headers Headers
	sign    func(req *http.Request, body any) string
}

type Headers map[string]string

func (c *HTTPClient) Send(uri, method string, data []byte) ([]byte, error) {
	var url string
	if c.baseURL == "" {
		url = uri
	} else {
		url = c.baseURL + uri
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, xerror.New(err.Error())
	}

	for key, value := range c.Headers {
		req.Header.Set(key, value)
	}

	if c.sign != nil {
		sign := c.sign(req, data)
		if sign != "" {
			req.Header.Set("sign", sign)
		}
	}

	xlog.Info("http", "url", url, "method", method, "data", string(data), "headers", req.Header)

	return do(req)
}

var defaultClient = http.Client{}

func do(req *http.Request) ([]byte, error) {
	client := defaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, xerror.New(err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, xerror.New(err.Error())
	}

	return body, nil
}

// Get 发送GET请求
func (c *HTTPClient) Get(endpoint string) ([]byte, error) {
	var url string
	if c.baseURL == "" {
		url = endpoint
	} else {
		url = c.baseURL + endpoint
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, xerror.New(err.Error())
	}

	// 设置请求头
	for key, value := range c.Headers {
		req.Header.Set(key, value)
	}
	return do(req)
}

// Post 发送POST请求
func (c *HTTPClient) Post(endpoint string, data []byte) ([]byte, error) {
	var url string
	if c.baseURL == "" {
		url = endpoint
	} else {
		url = c.baseURL + endpoint
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, xerror.New(err.Error())
	}

	// 设置请求头
	for key, value := range c.Headers {
		req.Header.Set(key, value)
	}

	return do(req)
}
