package util

import (
	"strconv"
	"time"
)

func Timestamp10(t ...time.Time) string {
	if len(t) == 0 {
		t = []time.Time{time.Now()}
	}
	return Itoa(int(t[0].Unix()))
}

func Timestamp13(t ...time.Time) string {
	if len(t) == 0 {
		t = []time.Time{time.Now()}
	}
	return Itoa(int(t[0].UnixMilli()))
}

func Itoa(num int) string {
	return strconv.FormatInt(int64(num), 10)
}

func Atoi(s string) int {
	num, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		panic(err)
	}
	return int(num)
}
