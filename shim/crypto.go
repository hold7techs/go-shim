package shim

import (
	"crypto/md5"
	"encoding/hex"
)

// ComputeMD5Hash 计算字符串的MD5哈希值
func ComputeMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)
}
