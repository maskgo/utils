# Util -- IP获取/转换

### 获取指定Prefix的IP地址

`注：prefix需根据当前所处网络设置指定`

```go
func Ip(prefix ...string) (string, error)
```

example:
```go
ipAddr, err := util.Ip("12.")
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println("ip: ", ipAddr)
}
```

### 转换

```go
func Ip2Long(ipAddr string) (uint32, error)
```

example:
```go
ip := "127.0.0.1"
ipInt, err := util.Ip2Long(ip)
if err != nil {
    fmt.Println(err)
} else{
    fmt.Println(ipInt)
}
```