package shim

import (
	"bufio"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
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

// ParseStrIDToUint 将strUID转成uin64
func ParseStrIDToUint[T uint32 | uint64 | uint](strUID string, dftUID T) T {
	parsedUID, err := strconv.ParseUint(strUID, 10, 64)
	if err != nil {
		return dftUID
	}
	return T(parsedUID)
}

// MustParseStrToTimeDuration 解析字符串格式时间
func MustParseStrToTimeDuration(str string) time.Duration {
	dur, err := time.ParseDuration(str)
	if err != nil {
		log.Fatalf("[MustParseStrToTimeDuration] %s", err)
	}

	return dur
}

// GetMapKeyValue 获取map m的key值，如果map为空或map中对应的key值不存在，则直接返回默认值T
// 注意: dft值类型不能设置错误，一定需要和map m值类型对齐
func GetMapKeyValue(m map[string]any, key string, dft any) any {
	// 检查map是否为空
	if m == nil {
		return dft
	}

	// 尝试从map中获取对应的key值
	value, exists := m[key]
	if !exists {
		return dft
	}

	// 返回默认值
	return value
}

const letters = "abcdefghijklmnopqrstuvwxyz0123456789"

// GenRandomLengthStr 获取随机长度的字符串
func GenRandomLengthStr(length int) string {
	// 创建一个字符切片来存放随机字符
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

// GenStreamStrChan 生成一个字符串流的 channel
func GenStreamStrChan(rawStr string, interval time.Duration) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out) // 确保在结束时关闭 channel
		for _, char := range rawStr {
			out <- string(char)  // 将每个字符发送到 channel
			time.Sleep(interval) // 等待指定的时间间隔
		}
	}()

	return out
}

// GenStreamFromReadFile 读取文件并生成字符串流的通道
func GenStreamFromReadFile(filename string, interval time.Duration) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out) // 确保在 goroutine 结束时关闭通道

		// 打开文件
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		// Scan文件
		scanner := bufio.NewScanner(file)

		// 设置分隔符为单个字符
		scanner.Split(bufio.ScanRunes)

		// 逐个字符扫描文件
		for scanner.Scan() {
			for _, char := range scanner.Text() {
				out <- string(char)
				time.Sleep(interval) // 等待指定的时间间隔
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
		}
	}()

	return out
}

// HashStringToUint64 将字符串哈希为指定显示长度的 uint64
func HashStringToUint64(s string, n int) uint64 {
	if n <= 0 || n > 64 {
		panic("n must be between 1 and 64")
	}

	// 创建 FNV-1a 哈希
	h := fnv.New64a()
	h.Write([]byte(s))

	// 获取哈希值
	hashValue := h.Sum64()

	// 计算最大值
	maxValue := uint64(math.Pow(2, float64(n))) - 1

	// 返回哈希值，限制在指定范围内
	return hashValue & maxValue
}
