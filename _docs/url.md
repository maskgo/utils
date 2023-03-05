# Util -- URL常用方法

### 解析URL

`component: -1: all; 1: scheme; 2: host; 4: port; 8: user; 16: pass; 32: path; 64: query; 128: fragment`
```go
func ParseUrl(str string, component int) (map[string]string, error)
```

example:
```go
url := "https://www.test.com"
scheme, err := util.ParseUrl(url)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(scheme)
}
```


### url 转义

```go
func UrlEncode(str string) string
```

example:
```go
url := "https://www.test.com/v1/s.s.get?uid=1233&name=小明"
fmt.Println(util.UrlEncode(url))
```

### url 解转义

```go
func UrlDecode(str string) (string, error)
```

example:
```go
rawUrl := "https%3A%2F%2Fwww.test.com%2Fv1%2Fs.s.get%3Fuid%3D1233%26name%3D%E5%B0%8F%E6%98%8E"
fmt.Println(util.UrlDecode(rawUrl))
```


