package log

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	std          = log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)
	defaultLevel = InfoLevel // 默认级别
)

func init() {
	// 基于环境变量，更新
	if logLevel := os.Getenv("LOG_LEVEL"); logLevel != "" {
		level, err := strconv.Atoi(logLevel)
		if err != nil {
			return
		}
		defaultLevel = Level(level)
	}
}

// Level log级别
type Level int

const (
	TraceLevel Level = iota
	DebugLevel
	WarnLevel
	InfoLevel
	ErrorLevel
	FatalLevel
)

var descMap = map[Level]string{
	TraceLevel: "TRACE",
	DebugLevel: "DEBUG",
	InfoLevel:  "INFO",
	WarnLevel:  "WARN",
	ErrorLevel: "ERROR",
	FatalLevel: "FATAL",
}

// Debug 调试日志输出
func Debug(v ...any) {
	printMessage(DebugLevel, v)
}

// Debugf 调试日志输出
func Debugf(format string, v ...any) {
	printFormatMessage(DebugLevel, format, v...)
}

func Info(v ...any) {
	printMessage(InfoLevel, v)
}

// Infof info级别消息格式化输出
func Infof(format string, v ...any) {
	printFormatMessage(InfoLevel, format, v...)
}

// Warn 错误信息
func Warn(v ...any) {
	printMessage(WarnLevel, v)
}

// Warnf 格式化输出
func Warnf(format string, v ...any) {
	printFormatMessage(WarnLevel, format, v...)
}

// Error 错误信息
func Error(v ...any) {
	printMessage(ErrorLevel, v)
}

// Errorf 格式化输出
func Errorf(format string, v ...any) {
	printFormatMessage(ErrorLevel, format, v...)
}

// Fatal 严重错误
func Fatal(v ...any) {
	printMessage(FatalLevel, v)
	os.Exit(1)
}

func Fatalf(format string, v ...any) {
	printFormatMessage(FatalLevel, format, v...)
	os.Exit(1)
}

func printMessage(level Level, v []any) {
	// 仅在log打印日志处调用的等级，大于系统配置等级时候才打印
	if level >= defaultLevel {
		std.Output(3, fmt.Sprintf("[%s] %s", descMap[level], v))
	}
}

func printFormatMessage(level Level, format string, v ...any) {
	// 仅在log打印日志处调用的等级，大于系统配置等级时候才打印
	if level >= defaultLevel {
		std.Output(3, fmt.Sprintf("[%s] %s", descMap[level], fmt.Sprintf(format, v...)))
	}
}
