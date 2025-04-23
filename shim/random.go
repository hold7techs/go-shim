package shim

import (
	"math/rand"
)

// RandElem 从给定的切片中随机选择一个元素
func RandElem[T any](list []T) T {
	if len(list) == 0 {
		var zero T
		return zero // 返回类型 T 的零值
	}
	index := rand.Intn(len(list)) // 生成一个随机索引
	return list[index]            // 返回随机选择的元素
}
