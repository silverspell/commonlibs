package commonlibs

import (
	"crypto/md5"
	"encoding/hex"
)

func Substring(in string, start, length int) string {
	size := len(in)
	endIndex := start + length

	if length == 0 || start < 0 || start > size || endIndex > size {
		return ""
	}

	return in[start:endIndex]
}

func Md5(in string) string {
	sum := md5.Sum([]byte(in))
	md5Bytes := sum[:]
	bytesOut := make([]byte, md5.Size*2)
	hex.Encode(bytesOut, md5Bytes)

	return string(bytesOut)
}
