# Util -- html预处理

### html 转义

```go
func Htmlentities(str string) string
```

example:
```go
const s = `"Fran & Freddie's Diner" <tasty@example.com>`
fmt.Println(util.Htmlentities(s))
```

### html 解转义
```go
func HtmlEntityDecode(str string) string
```

example:
```go
const s = `&quot;Fran &amp; Freddie&#39;s Diner&quot; &lt;tasty@example.com&gt;`
fmt.Println(util.HtmlEntityDecode(s))
```
