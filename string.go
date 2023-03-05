package utils

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

// 驼峰转换为下划线
func ToSnakeCase(str string, toLower bool) string {
	snake := matchAllCap.ReplaceAllString(matchFirstCap.ReplaceAllString(str, "${1}_${2}"), "${1}_${2}")

	if toLower {
		return strings.ToLower(snake)
	}

	return snake
}

// -/_ 转驼峰
func ToCamelCase(str string, lcFirst bool) string {
	data := make([]byte, 0, len(str))
	j, k, n := false, false, len(str)-1

	for i := 0; i <= n; i++ {
		d := str[i]
		if !k && d >= 'A' && d <= 'Z' {
			k = true
		}

		if d >= 'a' && d <= 'z' && (j || !k) {
			d -= 32
			j = false
			k = true
		}

		if k && (d == '_' || d == '-') && n > i && str[i+1] >= 'a' && str[i+1] <= 'z' {
			j = true
			continue
		}

		data = append(data, d)
	}

	camel := string(data)
	if lcFirst {
		camel = LcFirst(camel)
	}

	return camel
}

// 首字母大写
func UcFirst(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}

// 首字母小写
func LcFirst(n string) string {
	return strings.ToLower(n[:1]) + n[1:]
}

func Substr(str string, start uint, length int) string {
	if length < -1 {
		return str
	}

	switch {
	case length == -1:
		return str[start:]
	case length == 0:
		return ""
	}

	end := int(start) + length
	if end > len(str) {
		end = len(str)
	}

	return str[start:end]
}

func Strrev(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func MbStrlen(str string) int {
	return utf8.RuneCountInString(str)
}

func Trim(s, cutset string) string {
	if cutset == "" {
		return strings.TrimSpace(s)
	}

	return strings.Trim(s, cutset)
}

func LTrim(s, cutset string) string {
	return strings.TrimLeft(s, cutset)
}

func RTrim(s, cutset string) string {
	return strings.TrimRight(s, cutset)
}

func Chr(ascii int) string {
	return string(ascii)
}

func Ord(char string) int {
	r, _ := utf8.DecodeRune([]byte(char))

	return int(r)
}
