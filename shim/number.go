package shim

import (
	"fmt"
	"strings"
)

// Negative 确保返回负数金额值
func Negative[T int | int32 | int64 | float32 | float64](money T) T {
	if money > 0 {
		return -money
	}
	return money
}

// Positive 确保返回正数金额值
func Positive[T int | int32 | int64 | float32 | float64](money T) T {
	if money < 0 {
		return -money
	}
	return money
}

// InElems 检测某个元素elem是否在切片elems中
func InElems[T comparable](elem T, elems []T) bool {
	for _, v := range elems {
		if elem == v {
			return true
		}
	}
	return false
}

// UniqElems 从elems中剔除重复元素返回
func UniqElems[T comparable](elems []T) []T {
	m := make(map[T]struct{})
	var retIDs []T
	for _, id := range elems {
		if _, ok := m[id]; !ok {
			retIDs = append(retIDs, id)
			m[id] = struct{}{}
		}
	}
	return retIDs
}

// PagingElems 从elems切片中，按每页size大小，获取page页的数据
func PagingElems[T interface{}](elems []T, page int, size int) []T {
	if len(elems) == 0 {
		return elems
	}
	if page <= 0 {
		page = 1
	}

	total := len(elems)
	start := (page - 1) * size
	if start > total {
		start = total
	}
	end := start + size

	if end > total {
		end = total
	}

	return elems[start:end]
}

// ShardingElems 对elems 按批次大小batchSize分片
func ShardingElems[T comparable](elems []T, batchSize int) (batches [][]T) {
	lens := len(elems)
	if batchSize <= 0 {
		batchSize = lens
	}

	batchCount := lens / batchSize
	for i := 0; i <= batchCount; i++ {
		start := i * batchSize
		end := start + batchSize
		if end > lens {
			end = lens
		}
		if len(elems[start:end]) == 0 {
			continue
		}
		batches = append(batches, elems[start:end])
	}

	return
}

// JoinElems 对elems进行sep字符拼接，比如SQl场景等
func JoinElems[T comparable](elems []T, sep string) string {
	var ss []string
	for _, id := range elems {
		ss = append(ss, fmt.Sprintf("%v", id))
	}
	return strings.Join(ss, sep)
}
