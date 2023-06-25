package log

import (
	"fmt"
	"log"
	"os"
)

var std = log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)

// LogLevel log级别
type LogLevel int

const (
	TraceLevel LogLevel = iota
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

var descMap = map[LogLevel]string{
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

func printMessage(level LogLevel, v ...any) {
	std.Output(3, fmt.Sprintf("[%s] %s", descMap[level], v))
}

func printFormatMessage(level LogLevel, format string, v ...any) {
	std.Output(3, fmt.Sprintf("[%s] %s", descMap[level], fmt.Sprintf(format, v...)))
}
