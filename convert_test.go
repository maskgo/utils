package utils

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func TestToInt(t *testing.T) {
	ui8 := []uint8{32, 50, 48, 56, 52, 53, 49, 49, 55}
	fmt.Println(string(ui8))
	i := ToInt(ui8)
	fmt.Println(i)
}

func TestToString(t *testing.T) {
	var i64 int64 = 18
	s := ToString(i64)
	log.Println(s)
	var ui8 uint = 18
	s = ToString(ui8)
	log.Println(s)
	var ui64 uint64 = 18
	s = ToString(ui64)
	log.Println(s)

	var f32 float32 = 18.19998
	s = ToString(f32)
	log.Println(s)
	var f64 float64 = 18.19998
	s = ToString(f64)
	log.Println(s)
}

func TestStringMapConvert(t *testing.T) {
	var tp map[string]string
	args := make(map[string]interface{})
	args["name"] = "Alice"
	fmt.Println(StringMapConvert(args, reflect.TypeOf(tp)))
}

func TestSliceInterfaceConvert(t *testing.T) {
	var tp []int
	args := []interface{}{1, 3, 5}
	fmt.Println(SliceInterfaceConvert(args, reflect.TypeOf(tp)))
}
