# Util -- 文件操作

### 检查文件是否存在

```
func FileExists(path string) (bool, error)
```

example:
```go
path := "/etc/localtime"
if beExists, err := util.FileExists(path); !beExists {
    fmt.Println(err)
}
```

### 拷贝文件

```go
func FileCopy(dest, source string) error
```

example:

```go
desPath := "/tmp/localtime"
srcPaht := "/etc/localtime"

if err := util.FileCopy(desPath, srcPath); err != nil {
    fmt.Println(err)
}
```

### 创建文件夹

```go
func Mkdir(s string) error
```

examle:
```go
err := util.Mkdir("/tmp/test")
if err != nil {
    fmt.Println(err)
}
```

### 目录是否可写
```go
func IsWritable(path string) (isWritable bool, err error)
```

example:
```go
if ok, err := util.IsWritable("/usr/local"); !ok {
    fmt.Println(err)
}
```

### 检测目录

```go
func IsDir(path string) (bool, error)
```

example:

```go
if ok, err := util.IsDir("/usr/local"); !ok {
    fmt.Println(err)
}
```

### 文件行数

```go
func FileLineNum(path string) (num int, err error)
```

example:
```go
if ok, err := util.FileLineNum("/etc/hosts"); !ok {
    fmt.Println(err)
}
```

### 逐行扫描文件，同时并回调内容和当前行号

```go
func FileScan(path string, f func(b []byte, line int) bool, maxLineLen ...int) error
```
example:

```go
func test(b []byte, lineNo int) bool {
    fmt.Println(b)
    if lineNo == 2 {
        return true
    }
}
err := util.FileScan("/etc/hosts", test, 5)
if err != nil {
    fmt.Println(err)
}
```

### 获取目录下的文件

```go
func ScanDir(path string) ([]string, error)
```

example:
```go
dir := "/usr/local/bin"
l, err := util.ScanDir(dir)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(l)
}
```

### 获取路径下相关信息

`options: -1: all; 1: dirname; 2: basename; 4: extension; 8: filename`
```go
func Pathinfo(path string, options int) map[string]string
```

example:
```go
info := util.Pathinfo("/usr/local/bin", 1)
fmt.Printf("dirname: %s", info["dirname"])
```

### 写文件

```go
func FilePutContents(filename string, data string, mode os.FileMode) error
```
FileMode is just a uint32, ref <https://golang.org/pkg/os/#FileMode>

example:
```go
if err := FilePutContents("/tmp/test", "hello world", 0666); err != nil {
    fmt.Println(err)
}
```

### 读文件
```go
func FileGetContents(filename string) (string, error)
```

example:
```go
content, err := util.FileGetContents("/etc/hosts")
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(content)
}
```

### 删文件

```go
func Unlink(filename string) error
```

example:
```go
if err := util.Unlink("/tmp/test"); err != nil {
    fmt.Println(err)
}
```