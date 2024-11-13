package util

import (
	"encoding/hex"
)

func Pointer[T any](d T) *T {
	return &d
}

func ByteString(s string) []byte {
	bs, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return bs
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
