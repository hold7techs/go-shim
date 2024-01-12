package redisx

import (
	"fmt"
	"time"
)

const (
	// TTLForever 没有过期时间
	TTLForever = 0
)

// RdKey get redis key by format and id slice
func RdKey(format string, id ...interface{}) string {
	return fmt.Sprintf(format, id...)
}

// ParseKeyDuration get redis key time from string
func ParseKeyDuration(s string) (time.Duration, error) {
	dur, err := time.ParseDuration(s)
	if err != nil {
		return 0, err
	}
	return dur, nil
}

// MustParseKeyDuration get redis key time expire, if parse got err, expire time return 0
func MustParseKeyDuration(s string) time.Duration {
	dur, err := time.ParseDuration(s)
	if err != nil {
		return 0
	}
	return dur
}

// KeyFunc get redis key from interface object, g and its field value
//
//	eg. kfn := func(obj interface{}) string { return fmt.sprintf("uid:%s", obj.(*User).uid)}
type KeyFunc func(obj interface{}) string

// Uint64ToStringRdKeys []uint64 to []string
func Uint64ToStringRdKeys(format string, ids []uint64) []string {
	var keys []string
	for _, id := range ids {
		keys = append(keys, fmt.Sprintf(format, id))
	}
	return keys
}
