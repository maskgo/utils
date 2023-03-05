# Util -- 正则验证手机号和邮箱

### check mobile

```go
func IsMobile(s interface{}) bool
```

example:

```go
tel := "13911112222"
if yes := util.IsMobile(tel); !yes {
    fmt.Println("invalid telephon")
}
```

### check mail


```go
func IsMail(s interface{}) bool
```

example:

```go
mail := "abc@qq.com"
if yes := util.IsMobile(tel); !yes {
    fmt.Println("invalid mail")
}
```