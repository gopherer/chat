package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

//Md5Encode 小写
func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	tempStr := h.Sum(nil)
	return hex.EncodeToString(tempStr)
}

//MD5Encode 大写
func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

//MakePassword 加密
func MakePassword(plainPwd string, salt string) string {
	return Md5Encode(plainPwd + salt)
}

//ValidPassword 解密
func ValidPassword(plainPwd string, salt string, password string) bool {
	return Md5Encode(plainPwd+salt) == password
}
