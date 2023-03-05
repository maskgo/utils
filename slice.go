package utils

import (
	"reflect"
	"strings"
)

// 去空行
func SliceStringFilter(items []string) []string {
	result := []string{}

	for _, text := range items {
		if strings.Trim(text, " ") != "" {
			result = append(result, text)
		}
	}

	return result
}

// 去重
func SliceStringUnique(s *[]string) {
	f := make(map[string]bool, len(*s))
	t := 0

	for i, v := range *s {
		if _, ok := f[v]; !ok {
			f[v] = true
			(*s)[t] = (*s)[i]
			t++
		}
	}

	*s = (*s)[:t]
}

func SliceStringIn(elt string, slice []string) bool {
	for _, v := range slice {
		if v == elt {
			return true
		}
	}

	return false
}

func SliceIntIn(elt int, slice []int) bool {
	for _, v := range slice {
		if v == elt {
			return true
		}
	}

	return false
}

func SliceIn(elt, slice interface{}) bool {
	v := reflect.Indirect(reflect.ValueOf(slice))

	for i := 0; i < v.Len(); i++ {
		if reflect.DeepEqual(v.Index(i).Interface(), elt) {
			return true
		}
	}

	return false
}

func SliceColumn(sm []map[string]interface{}, field string) []interface{} {
	var res []interface{}
	for _, v := range sm {
		res = append(res, v[field])
	}

	return res
}

func SliceColumnInt64(sm []map[string]interface{}, field string) []int64 {
	var res []int64
	for _, v := range sm {
		res = append(res, v[field].(int64))
	}

	return res
}

func SliceStringDiff(in1, in2 []string) (out1, out2 []string) {
	for _, h := range in1 {
		if SliceStringIn(h, in2) {
			continue
		}

		out1 = append(out1, h)
	}

	for _, h := range in2 {
		if SliceStringIn(h, in1) {
			continue
		}

		out2 = append(out2, h)
	}

	return
}
