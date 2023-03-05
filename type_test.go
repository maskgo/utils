package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	s := "hello world"
	fmt.Println(IsEmpty(reflect.ValueOf(s)))
}
