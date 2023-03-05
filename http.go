package utils

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// Http请求
type HTTPRequestOption struct {
	Rawurl  string        // 请求地址
	Method  string        // 请求方法
	Host    string        // 指定Host
	Timeout time.Duration // 超时时间
}

func HTTPRequest(hro *HTTPRequestOption) (res []byte, err error) {
	u, err := url.Parse(hro.Rawurl)
	if err != nil {
		return res, err
	}

	rHost := u.Host

	if len(hro.Host) > 0 {
		u.Host = hro.Host
	}

	timeout := 1000 * time.Millisecond // 默认1秒超时
	if hro.Timeout > 0 {
		timeout = hro.Timeout
	}

	ctx, cancel := context.WithCancel(context.TODO())

	time.AfterFunc(timeout, func() {
		cancel()
	})

	req, err := http.NewRequest(hro.Method, u.String(), nil)
	if err != nil {
		return res, err
	}

	req.Host = rHost
	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return res, err
	}

	defer resp.Body.Close()

	res, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}

	return res, nil
}
