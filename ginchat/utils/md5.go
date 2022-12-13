package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func Md5Encode(data string) string {
	h := md5.New()
	_, err := h.Write([]byte(data))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return hex.EncodeToString(h.Sum(nil))
}

// 转大写
func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

// 加密
func MakePassword(plainpwd, salt string) string {
	return Md5Encode(plainpwd + salt)
}

// 解密
func ValidPassword(plainpwd, salt, password string) bool {
	return Md5Encode(plainpwd+salt) == password
}
