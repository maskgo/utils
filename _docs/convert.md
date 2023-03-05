# Util -- 数值转换

### 任意类型转int64

```go
func ToInt64(i interface{}) (v int64)
```

example:
```go
fmt.Println(util.ToInt64("64"))
```

### 任意类型转int

```go
func ToInt(i interface{}) (v int)
```

example:
```go
fmt.Println(util.ToInt("64"))
```

### 任意类型转string

```go
func ToString(i interface{}) (v string)
```

example:
```go
fmt.Println(util.ToInt(64))
```

### 任意类型转fload64

```go
func ToFloat64(i interface{}) (v float64)
```

example:
```go
fmt.Println(util.ToFloat64(64.5))
```

### []uint8 转字符串

```go
func Ui8ToA(i interface{}) string
```

example:
```go
u8 := []uint8{98, 99}
fmt.Println(util.Ui8ToA(u8))
```

### []uint8 转字符串字节

```go
func Ui8ToB(i interface{}) (b []byte)
```

example:
```go
u8 := []uint8{98, 99}
fmt.Println(util.Ui8ToB(u8))
```

### StringMapConvert

```go
func StringMapConvert(fs map[string]interface{}, ts reflect.Type) interface{}
```

example:
```go

```

### 任意类型转bool

```go
func ToBool(i interface{}) bool
```

example:
```go
fmt.Println(util.ToBool("true"))
fmt.Println(util.ToBool(1))
fmt.Println(util.ToBool(true))
```

### map[string]interface{} 转成约定的类型

```go
func StringMapConvert(fs map[string]interface{}, ts reflect.Type) interface{}
```

example:
```go
mm := make(map[string]interface{})
mm["name"] = "Alice"
fmt.Println(util.StringMapConvert(mm, reflect.ValueOf(mm)))
```

### []interface{} 转成约定的类型

```go
func SliceInterfaceConvert(fs []interface{}, ts reflect.Type) interface{}
```

example:
```go
var tp []int
args := []interface{}{1, 3, 5}
fmt.Println(util.SliceInterfaceConvert(args, reflect.TypeOf(tp)))
```
