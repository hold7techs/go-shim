package go_shim

import (
	"encoding/json"
	"fmt"
	"log"
)

// Boom 不应该进入的位置，测试调试环境避免意外情况
func Boom(remark string) {
	log.Fatalf("[Boom] - [%s] bad access!", remark)
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
