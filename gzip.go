package util

import (
	"bytes"
	"compress/gzip"
	"io"
)

func GzipCompress(data []byte) []byte {
	var buf bytes.Buffer
	writer := gzip.NewWriter(&buf)
	if _, err := writer.Write(data); err != nil {
		panic(err)
	}
	if err := writer.Close(); err != nil {
		panic(err)
	}
	return buf.Bytes()
}

func GzipUncompress(data []byte) []byte {
	var buf bytes.Buffer
	reader, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(&buf, reader); err != nil {
		panic(err)
	}
	if err := reader.Close(); err != nil {
		panic(err)
	}
	return buf.Bytes()
}
