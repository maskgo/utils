# Util -- 方法糖

### 三元运算

```go
func If(cond bool, v1, v2 interface{}) interface{}
```

example:
```go
type Person struct {
    Name string
    Age  int
}
p1 := Person{"Alice", 20}
P2 := Person{"Bob", 23}
leader := util.If(p1.Age > p2.Age, p1, p2)
fmt.Printf("leader is %s\n", leader.Name)
```

### map中检测key

```go
func IsSet(d map[string]interface{}, s string) bool
```

example:
```go
members := make(map[string]interface{})
members["Alice"] = 20
members["Bob"] = 21
if yes := util.IsSet(&members, "Steve"); !yes {
    fmt.Println("Steve is not in list")
}
```

### struct 转成 map

```go
func Struct2Map(obj interface{}) (m map[string]interface{})
```

example:
```go
type Person struct {
    Name     string
    Age      int
    Country  string
}
p := Person{"Alice", 20, "CHINA"}
pInfo := util.Struct2Map(p)
fmt.Println(pInfo)
```

