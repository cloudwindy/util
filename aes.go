package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

func AESCBCEncrypt(plaintext, key, iv []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	padded := pkcs5Padding([]byte(plaintext), block.BlockSize())
	cipherText := make([]byte, len(padded))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText, padded)
	return cipherText
}

func AESCBCDecrypt(ciphertext, key, iv []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	plaintext := make([]byte, len(ciphertext))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintext, ciphertext)
	return pkcs5Unpadding(plaintext)
}

func pkcs5Padding(data []byte, blockSize int) []byte {
	padding := (blockSize - len(data)%blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

func pkcs5Unpadding(data []byte) []byte {
	dataLen := len(data)
	unpadding := int(data[dataLen-1])
	return data[:(dataLen - unpadding)]
}
