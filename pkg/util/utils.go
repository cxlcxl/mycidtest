package util

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
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

func Offset(page, pageSize int) int {
	return (page - 1) * pageSize
}

func RandString(n int) string {
	randBytes := make([]byte, n/2)
	rand.Read(randBytes)
	return fmt.Sprintf("%x", randBytes)
}

func Sha1(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func Mkdir(p string, times int) {
	if times > 3 {
		return
	}
	_, err := os.Stat(p)
	if err == nil {
		return
	}
	if os.IsNotExist(err) {
		err = os.MkdirAll(p, os.ModePerm)
		if err == nil {
			return
		}
	}
	times++
	Mkdir(p, times)
}

func ArrayUnique[T string | int | int64 | int8 | uint | uint64 | uint8](arr []T) []T {
	var newArr []T
	tmp := make(map[T]struct{})
	for _, item := range arr {
		if _, ok := tmp[item]; !ok {
			newArr = append(newArr, item)
		}
		tmp[item] = struct{}{}
	}
	return newArr
}
