package shim

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

// Md5CheckSum 计算字符串的MD5哈希值
func Md5CheckSum(text string) string {
	h := md5.New()
	h.Write([]byte(text))
	checksum := h.Sum(nil)
	return hex.EncodeToString(checksum)
}

// Sha256Signature 使用sha256算法对数据message使用secret秘钥进行摘要签名
func Sha256Signature(secret string, message []byte) string {
	key := []byte(secret)

	h := hmac.New(sha256.New, key)
	h.Write(message)
	mac := h.Sum(nil)

	return hex.EncodeToString(mac)
}
