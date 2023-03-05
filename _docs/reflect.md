# Util -- 反射解析函数参数

### 参数反射解析

```go
func ReflectMethodArgsConvert(l int, m reflect.Value, data []interface{}) []reflect.Value
```

example:
```go
import "reflect"
import "fmt"
print := func(name string, age int, country string, l []int) {
    fmt.Printf("%s is %d, from %s, %v\n", name, age, country, l)
    return
}
args := []interface{}{"Alice", 18, "China", []interface{1, 3, 5}}
values := ReflectMethodArgsConvert(4, reflect.ValueOf(print), args)
for _, item := range values {
    fmt.Println(item.Interface())
}
```