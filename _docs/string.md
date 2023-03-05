# Util -- string

### camel case to snake case

```go
func ToSnakeCase(str string, toLower bool) string
```

example:
```go
s := "MyStruct"
fmt.Println(util.ToSnakeCase(s, true))
```

### snake case to camel case

```go
func ToCamelCase(str string, lcFirst bool) string
```

example:
```go
s := "my_struct"
fmt.Println(util.ToCamelCase(s, true))
```


### 首字母大写

```go
func Ucfirst(s string) string
```

example:
```go
s := "mystruct"
fmt.Println(util.Ucfirst(s))
```

### 首字母小写

```go
func LcFirst(s string) string
```

example:
```go
s := "MyStruct"
fmt.Println(util.LcFirst(s))
```

### 子串

```go
func Substr(str string, start uint, length int) string
```

example:
```go
s := "MyStruct"
fmt.Println(util.Substr(s, 2, 2))
```


### 反转

```go
func Strrev(str string) string 
```

example:
```go
s := "MyStruct"
fmt.Println(util.Strrev(s))
```


### 字符个数

```go
func MbStrlen(str string) int
```

example:
```go
s := "MyStruct"
fmt.Println(util.MbStrlen(s))
```

### 去掉指定字符串

```go
func Trim(s, cutset string) string
```

example:
```go
s := "MyStruct"
fmt.Println(util.Trim(s))
```


### 去掉左边开头的指定字符串

```go
func Ltrim(s, cutset string) string
```

example:
```go
s := " Hello world"
fmt.Println(util.Ltrim(s, "Helo"))
```

### ascii 转字符串

```go
func Chr(ascii int) string
```

example:
```go
fmt.Println(util.Chr(96))
```

### utf-8编码的码值的个数

```go
func Ord(char string) int
```

example:
```go
s := "Hello world"
fmt.Println(util.Ord(s))
```


