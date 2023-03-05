package utils

import (
	"fmt"
	"testing"
)

func TestRound(t *testing.T) {
	fmt.Println(Round(3.6))
}

func TestCeil(t *testing.T) {
	fmt.Println(Ceil(3.6))
}

func TestFloor(t *testing.T) {
	fmt.Println(Floor(3.6))
}

func TestMax(t *testing.T) {
	fmt.Println(Max(1.2, 4.9, 3.7, 1.1))
}

func TestMin(t *testing.T) {
	fmt.Println(Min(1.2, 4.9, 3.7, 1.1))
}
