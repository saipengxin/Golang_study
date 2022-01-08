package mylog

import (
	"fmt"
	"path"
	"runtime"
	"time"
)

// Logger 日志结构体，开头字母大写，其他包可以调用
type Logger struct{}

// NewLogger 构造函数
func NewLogger() *Logger {
	return &Logger{}
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

// Debug Debug级别日志
func (l Logger) Debug(msg string) {
	now := time.Now().Format("2006-01-02 15:04:05")
	funcName, fileName, lineNo := getInfo(2)
	fmt.Printf("[%s] [Debug] [%s:%s:%d] %s\n", now, fileName, funcName, lineNo, msg)
}

// Trace Trace级别日志
func (l Logger) Trace(msg string) {
	now := time.Now().Format("2006-01-02 15:04:05")
	funcName, fileName, lineNo := getInfo(2)
	fmt.Printf("[%s] [Trace] [%s:%s:%d] %s\n", now, fileName, funcName, lineNo, msg)
}

// Info Info级别日志
func (l Logger) Info(msg string) {
	now := time.Now().Format("2006-01-02 15:04:05")
	funcName, fileName, lineNo := getInfo(2)
	fmt.Printf("[%s] [Info] [%s:%s:%d] %s\n", now, fileName, funcName, lineNo, msg)
}

// Warning Warning级别日志
func (l Logger) Warning(msg string) {
	now := time.Now().Format("2006-01-02 15:04:05")
	funcName, fileName, lineNo := getInfo(2)
	fmt.Printf("[%s] [Warning] [%s:%s:%d] %s\n", now, fileName, funcName, lineNo, msg)
}

// Error Error级别日志
func (l Logger) Error(msg string) {
	now := time.Now().Format("2006-01-02 15:04:05")
	funcName, fileName, lineNo := getInfo(2)
	fmt.Printf("[%s] [Error] [%s:%s:%d] %s\n", now, fileName, funcName, lineNo, msg)
}

// Fatal Fatal级别日志
func (l Logger) Fatal(msg string) {
	now := time.Now().Format("2006-01-02 15:04:05")
	funcName, fileName, lineNo := getInfo(2)
	fmt.Printf("[%s] [Fatal] [%s:%s:%d] %s\n", now, fileName, funcName, lineNo, msg)
}
