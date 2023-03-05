package utils

import "testing"

var testColorMsg = "hello maskwang"

func TestRandomColor(t *testing.T) {
	t.Log(RandomColor())
}

func TestYellow(t *testing.T) {
	t.Log(Yellow(testColorMsg))
}

func TestRed(t *testing.T) {
	t.Log(Red(testColorMsg))
}

func TestRedf(t *testing.T) {
	t.Log(Redf(testColorMsg, "kaka"))
}

func TestBlue(t *testing.T) {
	t.Log(Blue(testColorMsg))
}

func TestGreen(t *testing.T) {
	t.Log(Green(testColorMsg))
}

func TestGreenf(t *testing.T) {
	t.Log(Greenf(testColorMsg, "kaka"))
}
