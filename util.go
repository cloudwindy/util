package util

import (
	"encoding/hex"
	"strconv"
)

func Pointer[T any](d T) *T {
	return &d
}

func ParseHexBytes(s string) []byte {
	bs, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return bs
}

func ParseHexInt(s string) int {
	n, err := strconv.ParseInt(s, 16, 0)
	if err != nil {
		panic(err)
	}
	return int(n)
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
