# Utils -- DeepCopy

### 切片深拷贝

util/bytes.go

```go
func BytesClone(b []byte) []byte
```

example:
```go
func TestBytesClone(t *testing.T) {
	slice1 := []byte{0, 2, 4, 6, 8}
	fmt.Printf("slice_1 mem addr:\t %p\n", &slice1)
	sliceCopy := BytesClone(slice1)
	fmt.Printf("sliceCopy memory addr:\t %p\n", &sliceCopy)
	fmt.Println(sliceCopy)
}
```