package shim

import (
	"encoding/json"
	"fmt"
	"log"
)

// Boom 不应该进入的位置，测试调试环境避免意外情况
func Boom(remark string) {
	log.Fatalf("[Boom] - [%s]  Met a Boom!!", remark)
}

// ToJsonString 将变量转成Json类型，主要for debug
func ToJsonString(v interface{}, pretty bool) string {
	if v == nil {
		return "<nil>"
	}
	var b []byte
	var err error
	if pretty {
		b, err = json.MarshalIndent(v, "", "  ")
	} else {
		b, err = json.Marshal(v)
	}
	if err != nil {
		return fmt.Sprintf("[error] can not marshal: %s", err)
	}
	return string(b)
}

// ProcessStringsSlice 返回值为自定义处理函数针对切片每个元素进行处理后，返回新的切片内容
func ProcessStringsSlice(strs []string, filter func(string) bool, fn func(string) string) []string {
	var result []string

	for _, str := range strs {
		// 过滤函数检测
		if filter != nil && filter(str) {
			continue
		}

		if fn != nil {
			str = fn(str)
		}
		result = append(result, str)
	}

	return result
}
