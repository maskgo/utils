package utils

import (
	"fmt"
	"testing"
)

func TestUrlEncode(t *testing.T) {
	url := "https://www.test.com/v1/s.s.get?uid=1233&name=小明"
	fmt.Println(URLEncode(url))
}
