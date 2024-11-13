package util

import (
	"crypto/md5"
	"crypto/sha256"
)

func SHA256Sum(data []byte) []byte {
	sum := sha256.Sum256(data)
	return sum[:]
}

func MD5Sum(data []byte) []byte {
	sum := md5.Sum(data)
	return sum[:]
}
