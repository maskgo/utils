# Util -- dump interface

### dump interface

```go
func Dump(v ...interface{}) string
```

example:
```go
type Person struct {
    Name    string
    Age     int
    Country string
}
p := Person{"Alice", 18, "China"}
fmt.Println(util.Dump(p))
```