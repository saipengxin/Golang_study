package mylog

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

type Logger interface {
	Debug(format string, a ...interface{})
	Trace(format string, a ...interface{})
	Info(format string, a ...interface{})
	Warning(format string, a ...interface{})
	Error(format string, a ...interface{})
	Fatal(format string, a ...interface{})
}

// NewLogger 返回Logger类型的数据，因为这Console和file都实现了这个接口，所以既可以返回Console，又可以返回file
func NewLogger(t, lv string) Logger {
	var l Logger
	t = strings.ToLower(t)
	switch t {
	case "console":
		l = NewConsoleLogger(lv)
	case "file":
		l = NewFileLogger(lv, "./", "demoLog.log", 10*1024*1024)
	}
	return l
}

// 自定义类型
type LogLevel uint16

// 将对应的等级定义成数值
const (
	UNKNOWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

// 定义函数，将传递进来的字符串解析成对应等级数值
func parseLogLevel(s string) (LogLevel, error) {
	// 将字符串统一变为小写比较
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNOWN, err
	}
}

// 定义函数将对应等级的数值转换成字符串
func getLogString(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	}
	return "DEBUG" // 默认就是Debug等级
}

// getInfo 获取调用者相关信息
func getInfo(skin int) (string, string, int) {
	pc, file, lineNo, ok := runtime.Caller(skin)
	if !ok {
		fmt.Println("错误")
	}
	// 获取调用的函数名
	funcName := runtime.FuncForPC(pc).Name()
	// 文件名
	fileName := path.Base(file)
	return funcName, fileName, lineNo
}
