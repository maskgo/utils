# Util -- MD5/SHA/CRC/BASE64

util/hash.go

### Md5 生成一个字符串的摘要

```
func Md5(str string) string
```

example:
```
s := "hello world"
md5 := util.Md5(s)
fmt.Printf("%s md5: %s", s, md5)
```

### Md5File 生成一个文件的摘要

```
func Md5File(path string) (string, error)
```

example:
```
path := "/etc/localtime"
fileHash, err := util.Md5File(path)
fmt.Printf("file %s hash : %v %v", path, fileHash, err)
```

### Sha1 为字符串生成160bit的摘要

```
func Sha1(str string) string 
```

example:
```
s := "hello world"
shaHex := util.Sha1(s)
fmt.Printf("%s sha1: %s", s, shaHex)
```

### Sha1File 生成文件的160bit摘要
```
func Sha1File(path string) (string, error)
```

example:
```
path := "/etc/localtime"
fileShaHex, err := util.Sha1File(path)
fmt.Printf("file %s hash : %v %v", path, fileHash, err)
```

### Crc16/Crc32 生成循环冗余校验码

```
func Crc32(str string) uint32 
func Crc16(str string) int
```

example:
```
s := "hello world"
scrc32 := util.Crc32(s)
scrc16 := util.Crc16(s)
fmt.Printf("%s crc32 %d, crc16 %d", s, scrc32, scrc16)
```

### Base64

```
func Base64Encode(str string) string
func Base64Decode(str string) (string, error)
```

example:
```
s := "hello world"
b64Hex := util.Base64Encode(s)
fmt.Printf("%s base64 encode: %s\n", s, b64Hex)

ss, err := util.Base64Decode(b64Hex)
fmt.Printf("%s base64 decode: %s", b64Hex, ss)
```