package utils

import "strconv"

// 数据库扫描类
type SRow struct {
	V interface{}
}

func (sr *SRow) Int64() int64 {
	return ToInt64(sr.V)
}

func (sr *SRow) Int() int {
	return ToInt(sr.V)
}

func (sr *SRow) Float64() float64 {
	f, _ := strconv.ParseFloat(UI8ToA(sr.V), 64)

	return f
}

func (sr *SRow) String() string {
	return UI8ToA(sr.V)
}

func (sr *SRow) Bytes() []byte {
	return UI8ToB(sr.V)
}
