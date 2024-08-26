package util

import (
	"crypto/md5"
	"encoding/hex"
)

func Password(password string, isMd5 bool) string {
	if !isMd5 {
		password = Md5(password)
	}
	return Md5(Md5(password))
}

func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
