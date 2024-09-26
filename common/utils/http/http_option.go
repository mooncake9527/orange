package httpUtil

import "net/http"

const (
	HeaderSignature   = "signature"
	HeaderContentType = "Content-Type"

	ContentTypeJson           = "application/json"
	ContentTypeXml            = "application/xml"
	ContentTypeFormUrlEncoded = "application/x-www-form-urlencoded"
)

type HTTPOption func(*httpOptions)

type httpOptions struct {
	baseUrl     string
	contentType string
	sign        func(req *http.Request, body any) string
}

func defaultHTTPOptions() *httpOptions {
	return &httpOptions{
		contentType: ContentTypeJson,
	}
}

func WithContentType(contentType string) HTTPOption {
	return func(o *httpOptions) {
		o.contentType = contentType
	}
}

func WithSign(sign func(req *http.Request, data any) string) HTTPOption {
	return func(o *httpOptions) {
		o.sign = sign
	}
}

func (o *httpOptions) apply(opts ...HTTPOption) {
	for _, opt := range opts {
		opt(o)
	}
}
