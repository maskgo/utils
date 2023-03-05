package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIf(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	p1 := Person{"Alice", 20}
	p2 := Person{"Bob", 23}
	leader := If(p1.Age > p2.Age, p1, p2)
	fmt.Println(leader)
}

func TestIsSet(t *testing.T) {
	members := make(map[string]interface{})
	members["Alice"] = 20
	members["Bob"] = 21
	if yes := IsSet(members, "Steve"); !yes {
		fmt.Println("Steve is not in list")
	}
}

func TestStruct2Map(t *testing.T) {
	type Object struct {
		Id   int
		Name string
	}

	object := Object{
		Id:   1,
		Name: "test",
	}

	m := Struct2Map(object)
	dm := map[string]interface{}{
		"Id":   1,
		"Name": "test",
	}

	if !reflect.DeepEqual(m, dm) {
		t.Error("not equal")
	}
}
