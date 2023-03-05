package utils

import "reflect"

func If(cond bool, v1, v2 interface{}) interface{} {
	if cond {
		return v1
	}

	return v2
}

func IsSet(d map[string]interface{}, s string) bool {
	if _, ok := d[s]; ok {
		return true
	}

	return false
}

func Struct2Map(obj interface{}) (m map[string]interface{}) {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	m = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		m[t.Field(i).Name] = v.Field(i).Interface()
	}

	return
}
