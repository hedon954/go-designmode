package main

import (
	"net/http"
	"strings"
	"time"
)

func main() {
	url := ""
	headers := make(map[string]string)
	timeout := time.Second
	data := ""
	HttpRequest("GET", url)
	HttpRequest("POST", url, WithHeaders(headers))
	HttpRequest("POST", url, WithTimeout(timeout), WithHeaders(headers), WithData(data))
}

// requestOption 针对可选的 HTTP 请求配置项
type requestOption struct {
	timeout time.Duration
	data    string
	headers map[string]string
}

type Option struct {
	apply func(*requestOption)
}

// defaultRequestOptions 默认请求选项
func defaultRequestOptions() *requestOption {
	return &requestOption{
		timeout: 5 * time.Second,
		data:    "",
		headers: nil,
	}
}

func WithTimeout(timeout time.Duration) *Option {
	return &Option{
		apply: func(option *requestOption) {
			option.timeout = timeout
		},
	}
}

func WithData(data string) *Option {
	return &Option{
		apply: func(option *requestOption) {
			option.data = data
		},
	}
}

func WithHeaders(headers map[string]string) *Option {
	return &Option{
		apply: func(option *requestOption) {
			option.headers = headers
		},
	}
}

// HttpRequest 省略返回值
func HttpRequest(method string, url string, options ...*Option) {
	reqOpts := defaultRequestOptions() // 默认的请求选项
	for _, opt := range options {      // 在reqOpts上应用通过options设置的选项
		opt.apply(reqOpts)
	}
	// 创建请求对象
	req, err := http.NewRequest(method, url, strings.NewReader(reqOpts.data))
	if err != nil {
		// 异常处理
	}

	// 设置请求头
	for key, value := range reqOpts.headers {
		req.Header.Add(key, value)
	}
	// 发起请求

	return
}
