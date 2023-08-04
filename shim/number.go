package shim

import (
	"fmt"
	"strings"
)

// Negative 确保金额类型为负数
func Negative[T int | int32 | int64 | float32 | float64](money T) T {
	if money > 0 {
		return -money
	}
	return money
}

// Positive 确保金额类型为正数
func Positive[T int | int32 | int64 | float32 | float64](money T) T {
	if money < 0 {
		return -money
	}
	return money
}

// InSlice 尝试用范型比较
func InSlice[T comparable](id T, ids []T) bool {
	for _, v := range ids {
		if id == v {
			return true
		}
	}
	return false
}

// Uniq 切片去重
func Uniq[T comparable](ids []T) []T {
	m := make(map[T]struct{})
	var retIDs []T
	for _, id := range ids {
		if _, ok := m[id]; !ok {
			retIDs = append(retIDs, id)
			m[id] = struct{}{}
		}
	}
	return retIDs
}

// PageData 获取分页数据
func PageData[T interface{}](data []T, page int, pageSize int) []T {
	if len(data) == 0 {
		return data
	}
	if page <= 0 {
		page = 1
	}

	total := len(data)
	start := (page - 1) * pageSize
	if start > total {
		start = total
	}
	end := start + pageSize

	if end > total {
		end = total
	}

	return data[start:end]
}

// ShardingNumbers 分片函数，按指定大小分成多片数据
func ShardingNumbers[T comparable](ids []T, batchSize int) (batches [][]T) {
	slen := len(ids)
	if batchSize <= 0 {
		batchSize = slen
	}

	batchCount := slen / batchSize
	for i := 0; i <= batchCount; i++ {
		start := i * batchSize
		end := start + batchSize
		if end > slen {
			end = slen
		}
		if len(ids[start:end]) == 0 {
			continue
		}
		batches = append(batches, ids[start:end])
	}

	return
}

// JoinNumbers 拼接数值成字符串
func JoinNumbers[T comparable](ids []T, sep string) string {
	var ss []string
	for _, id := range ids {
		ss = append(ss, fmt.Sprintf("%v", id))
	}
	return strings.Join(ss, sep)
}
