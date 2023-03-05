package utils

import (
	"fmt"
	"testing"
)

func TestRand(t *testing.T) {
	fmt.Println(Rand(123, 345))
}

func TestRandString(t *testing.T) {
	fmt.Println(RandString(10))
}
