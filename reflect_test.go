package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflectMethodArgsConvert(t *testing.T) {
	print := func(name string, age int, country string, l []int, m map[string]int) {
		fmt.Printf("%s is %d, from %s, %v\n", name, age, country, l)
		return
	}
	mm := make(map[string]interface{})
	mm["tel"] = 11
	args := []interface{}{"Alice", 18, "China", []interface{}{1, 3, 5}, mm}
	values := ReflectMethodArgsConvert(5, reflect.ValueOf(print), args)
	for _, item := range values {
		fmt.Println(item.Interface())
	}
}
