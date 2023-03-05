# Util -- 空值判定

### 检测各类型值是否为空

```go
func IsEmpty(val reflect.Value) bool
```

example:

```go
import reflect
s := "hello world"
fmt.Println(util.IsEmpty(reflect.ValueOf(s)))

m := make(map[string]int)
fmt.Println(util.IsEmpty(reflect.ValueOf(m)))
```
