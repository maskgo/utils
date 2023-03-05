# Util -- Json序列化和反序列化

### 序列化
```go
func JsonEncode(data interface{}) ([]byte, error)
```

example:
```go
type Person struct {
    Age  int
    Name string
}
person := Person {18, "Alice"}
b, err := util.JsonEncode(person)
if err != nil {
    fmt.Println(err)
}
fmt.Printf("person : %s\n", string(b))
```

### 反序列化

```go
func JsonDecode(data []byte, v interface{}) error
```

example:

```go
type Person struct {
    Age  int
    Name string
}
b := []byte(`"{Age:18,Name:Alice}"`)
person := Person{}
err := util.JsonDecode(b, &person)
if err != nil {
    fmt.Println(err)
}
```