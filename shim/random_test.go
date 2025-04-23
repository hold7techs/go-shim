package shim

import (
	"testing"
)

func TestRandElem(t *testing.T) {

	// 示例：从字符串切片中随机选择一个元素
	strings := []string{"apple", "banana", "cherry", "date"}
	randomString := RandElem(strings)
	t.Logf("随机选择的字符串: %s", randomString)

	// 示例：从整数切片中随机选择一个元素
	integers := []int{1, 2, 3, 4, 5}
	randomInt := RandElem(integers)
	t.Logf("随机选择的字符串: %v", randomInt)
}
