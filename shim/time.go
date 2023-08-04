package shim

import (
	"fmt"
	"time"
)

// GetTimeVersion 获取时间版本信息
func GetTimeVersion() string {
	t := time.Now()
	return fmt.Sprintf("v%s_%d", t.Format("20060102150405"), t.Nanosecond())
}

const (
	StdDataLayout            = `2006-01-02`
	StdDateTimeLayout        = `2006-01-02 15:04:05`
	StdCompactDateLayout     = `20060102`
	StdCompactDateTimeLayout = `20060102150405`
)

// StdDateStr 标准日期字符串
type StdDateStr string

// GetTime 日期字符串转时间
func (s StdDateStr) GetTime() time.Time {
	parseTime, err := time.ParseInLocation(StdDataLayout, string(s), time.Local)
	if err != nil {
		return time.Time{}
	}
	return parseTime
}

// StdDateTimeStr 标准日期时间格式字符串
type StdDateTimeStr string

// GetTime 日期时间字符串转时间
func (s StdDateTimeStr) GetTime() time.Time {
	parseTime, err := time.ParseInLocation(StdDateTimeLayout, string(s), time.Local)
	if err != nil {
		return time.Time{}
	}
	return parseTime
}

// TimestampToLayout 整型时间戳转成指定类型的时间字符串
func TimestampToLayout(timestamp int64, layout string) string {
	t := time.Unix(timestamp, 0)
	return t.Format(layout)
}
