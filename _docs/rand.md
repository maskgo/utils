# Util -- 随机数

### 指定区间的随机数

```go
func Rand(min int, max int) (r int)
```

example:
```go
fmt.Println(util.Rand(20, 100))
```

### 长度n的随机字符串

```go
func RandString(n int) string
```

example:
```go
fmt.Println(util.RandString(10))
```



