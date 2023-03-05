package utils

import (
	"fmt"
	"testing"
)

func TestDump(t *testing.T) {
	type Person struct {
		Name    string
		Age     int
		Country string
	}
	p := Person{"Alice", 18, "China"}
	fmt.Println(Dump(p))
}
