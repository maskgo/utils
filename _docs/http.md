# Util -- http

### 发起http请求

```go
type HttpRequestOption struct {
    Rawurl  string        // 请求地址
    Method  string        // 请求方法
    Host    string        // 指定Host
    Timeout time.Duration // 超时时间
}
func HttpRequest(hro *HttpRequestOption) (res []byte, err error)
```

example:
```go
hr := util.HttpRequestOption{
    Rawurl: "https://www.test.com"
    Method: "GET"
    Host:   "www.test.com"
    Timeout: time.Duration(3)
}
if resp, err := util.HttpRequest(&hr); err != nil {
    fmt.Println(err)
}else {
    fmt.Println(string(res))
}
```