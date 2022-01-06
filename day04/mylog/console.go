package mylog

import (
	"fmt"
	"time"
)

// ConsoleLogger 控制台日志结构体，开头字母大写，其他包可以调用，命名为ConsoleLogger是为了和文件日志的结构体区分
type ConsoleLogger struct {
	Lever LogLevel
}

// NewConsoleLogger 构造函数,接收一个日志等级作为参数
func NewConsoleLogger(levelStr string) *ConsoleLogger {
	// 将字符串日志等级转为数值型，方便比较
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return &ConsoleLogger{
		Lever: level,
	}
}

// 定义函数，用来判断当前日志等级是否允许输出
func (c ConsoleLogger) enable(logLevel LogLevel) bool {
	return logLevel >= c.Lever
}

// 不同级别的日志中，输出内容的部分重复性太高，这里提取出来封装成方法
// 这里定义成方法，如果是函数就和file文件中的函数重名了，同属一个包
func (c ConsoleLogger) log(lv LogLevel, format string, a ...interface{}) {
	if c.enable(lv) {
		// 字符串和参数拼接在一起才是完整的msg信息，我们这里直接使用原生的Sprintf来进行替换。
		msg := fmt.Sprintf(format, a...)
		now := time.Now().Format("2006-01-02 15:04:05")
		funcName, fileName, lineNo := getInfo(3)
		// getLogString 将日志级别数值转换成字符串
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now, getLogString(lv), fileName, funcName, lineNo, msg)
	}
}

// Debug Debug级别日志,第一个是一个字符串，第二个是任意类型，数量不定的参数
func (c ConsoleLogger) Debug(format string, a ...interface{}) {
	c.log(DEBUG, format, a...)
}

// Trace Trace级别日志
func (c ConsoleLogger) Trace(format string, a ...interface{}) {
	c.log(TRACE, format, a...)
}

// Info Info级别日志
func (c ConsoleLogger) Info(format string, a ...interface{}) {
	c.log(INFO, format, a...)
}

// Warning Warning级别日志
func (c ConsoleLogger) Warning(format string, a ...interface{}) {
	c.log(WARNING, format, a...)
}

// Error Error级别日志
func (c ConsoleLogger) Error(format string, a ...interface{}) {
	c.log(ERROR, format, a...)
}

// Fatal Fatal级别日志
func (c ConsoleLogger) Fatal(format string, a ...interface{}) {
	c.log(FATAL, format, a...)
}
