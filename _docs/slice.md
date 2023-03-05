# Util -- slice

### 去空格

```go
func SliceStringFilter(items []string) []string
```

example:
```go
ts := []string{
    "hello world",
    "Bye Bye"
}
fmt.Println(util.SliceStringFilter(ts))
```

### 去重

```go
func SliceStringUnique(s *[]string)
```

example:
```go
ts := []string{
    "hello",
    "Bye Bye"，
    "hello"
}
util.SliceStringUnique(&ts)
fmt.Println(ts)
```

### check string

```go
func SliceStringIn(elt string, slice *[]string) bool
```

example:
```go
ts := []string{
    "hello",
    "Bye Bye"，
    "hello"
}
if ok := util.SliceStringIn("hello", &ts); !ok {
    fmt.Println("hello is not exist")
}
```

### check int

```go
func SliceIntIn(elt int, slice *[]int) bool
```

example:
```go
ti := []int{1, 3, 6, 0}
if ok := util.SliceIntIn(5, &ti); !ok {
    fmt.Println("5 is not exist")
}
```

### check interface

```go
func SliceIn(elt, slice interface{}) bool
```

example:
```go
type Person struct {
    Age  int
    Name string
}
p1 := Person{18, "Alice"}
p2 := Person{20, "Bob"}
pl := []Person{p1, p2}
if ok := util.SliceIn(p1, pl); !ok {
    fmt.Println("p1 is not exist")
}
```

### 获取指定字段的所有值(interface)

```go
func SliceColumn(sm *[]map[string]interface{}, field string) []interface{}
```

example:
```go
type Person struct {
    Age     int
    Country string
}
p1 := Person{18, "China"}
p2 := Person{20, "USA"}
m := make(map[string]interface{})
m["Alice"] = p1
m["Bob"] = p2
countries := util.SliceColumn(&m, "Country")
fmt.Println(countries)
```

### 获取指定字段的所有值(int)

```go
func SliceColumnInt64(sm *[]map[string]interface{}, field string) []int64
```

example:
```go
type Person struct {
    Age     int
    Country string
}
p1 := Person{18, "China"}
p2 := Person{20, "USA"}
m := make(map[string]interface{})
m["Alice"] = p1
m["Bob"] = p2
countries := util.SliceColumnInt64(&m, "Age")
fmt.Println(countries)
```



