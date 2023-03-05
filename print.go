package utils

import (
	"github.com/liudng/godump"
)

func Dump(v ...interface{}) string {
	l := len(v)
	s := false

	if l > 1 {
		if v, ok := v[l-1].(bool); ok && v {
			s = true
			l--
		}
	}

	r := ""

	for i := 0; i < l; i++ {
		if s {
			r += godump.Sdump(v[i])
		} else {
			godump.Dump(v[i])
		}
	}

	return r
}
