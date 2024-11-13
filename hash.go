package util

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

func SHA256Sum(data []byte) string {
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:])
}

func MD5Sum(data []byte) string {
	sum := md5.Sum(data)
	return hex.EncodeToString(sum[:])
}
