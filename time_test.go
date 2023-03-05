package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestLastDates(t *testing.T) {
	ds := LastDates(5)   // 当前日期:20180417
	fmt.Println(ds)      // [20180416 20180415 20180414 20180413 20180412]
	ds = LastDates(5, 3) // 当前日期:20180417
	fmt.Println(ds)      // [20180414 20180413 20180412]
}

func TestExpireTimeByDay(t *testing.T) {
	n := time.Now().Unix()         // 1523952188
	d := Date("Y-m-d H:i:s", n)    // 2018-04-17 16:03:08
	s := ExpireTimeByDay(3)        // 201411
	nd := Date("Y-m-d H:i:s", n+s) // 2018-04-19 23:59:59

	fmt.Println(n, d, s, nd)
}

func TestDate(t *testing.T) {
	td := int(1524198841)
	ti64 := int64(td)
	tt := time.Unix(ti64, 0)
	ti := int(tt.Unix())

	v := "2018-04-20 12:34:00"
	if vv := Date("Y-m-d H:i:s", ti64); v != vv {
		t.Error(vv)
	}
	if vv := Date("Y-m-d H:i:s", ti); v != vv {
		t.Error(vv)
	}
	if vv := Date("Y-m-d H:i:s", tt); v != vv {
		t.Error(vv)
	}
}

func TestTimestampBeforeDay(t *testing.T) {

	ts := TimestampBeforeDay(6)
	fmt.Println(ts)

	d := Date("Y-m-d", ts)
	fmt.Println(d)

}

func TestExecTime(t *testing.T) {
	fmt.Println(ExecTime(func() string {
		return fmt.Sprintln("testaaa")
	}, false))

	fmt.Println(ExecTime(func() string {
		return fmt.Sprintln("testbbb")
	}, true))
}

func TestAvgTime(t *testing.T) {
	times := []time.Duration{20, 60, 50}
	fmt.Println(AvgTime(times))
}
