# Util -- 时间

### 获取前n天日期


```go
func LastDates(n int, es ...int) (dates []string)
```

example:
```go
// 前5天的日期
ds := util.LastDates(5)
fmt.Println(ds)
// 前5天中的前3天日期
ds = util.LastDates(5, 3)
fmt.Println(ds)
```

### 格式化

```go
func Date(format string, i ...interface{}) string
```

example:
```go
t64 := int64(1524198841)
format := "Y-m-d H:i:s"
fmt.Println(util.Date(format, t64))
```

### 字符串类型转time.Time

```go
func StrToTime(value string) (time.Time, error)
```

example:
```go
d := "2019-08-22"
if _, err := util.StrToTime(d); err != nil {
    fmt.Println(err)
}
```

### 时间戳（秒/毫秒/微秒）

`v[0]: 10--秒, 13--毫秒, 16--微秒`
```go
func Timestamp(v ...interface{}) int64
```

example:
```go
timestamp := util.Timestamp(13, time.Now())
fmt.Println(timestamp)
```

### 获取距离N天之后的秒数

```go
func ExpireTimeByDay(d ...int) (second int64)
```

example:
```go
s := ExpireTimeByDay(3)
fmt.Println(s)
```

### 获取回调函数执行时间

```go
func ExecTime(f func() string, d ...bool) (content string)
```

example:
```go
fmt.Println(util.ExecTime(func () string {
    return fmt.Sprintln("test")
}, false))

fmt.Println(util.ExecTime(func () {
    return fmt.Sprintln("test")
}, true))
```

### 平均时间

```go
func AvgTime(items []time.Duration) time.Duration
```

example:
```go
times := []time.Duration{20, 60, 50}
fmt.Println(util.AvgTime(times))
```

### 获取几天前的开始时间戳

```go
func TimestampBeforeDay(b int) int64
```

example:
```go
fmt.Println(util.TimestampBeforeDay(4))
```


### sleep(秒)

```go
func Sleep(t int64)
```

example:
```go
util.Sleep(3600)
```

### usleep(微秒)

```go
func Usleep(t int64)
```

example:
```go
util.Usleep(3600)
```
