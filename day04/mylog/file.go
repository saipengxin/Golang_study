package mylog

import (
	"fmt"
	"os"
	"path"
	"time"
)

// FileLogger 日志结构体，开头字母大写，其他包可以调用
type FileLogger struct {
	Lever       LogLevel // LogLevel 直接调用公共文件中的类型，同属于一个包，开头大写，可以调用
	filePath    string   // 日志文件保存的路径
	fileName    string   // 日志文件保存的文件名
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64 // 文件切割使用
}

// NewFileLogger 构造函数,接收一个日志等级作为参数
func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	// 将字符串日志等级转为数值型，方便比较
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		Lever:       level,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
	// 我们要在初始化的是否就将文件给准备好
	err = fl.initFile()
	if err != nil {
		panic(err)
	}
	return fl
}

// 不用传递参数，接收者直接就携带过来了
func (f *FileLogger) initFile() error {
	// 将文件路径和文件名称拼接成完整的路径
	//Join函数可以将任意数量的路径元素放入一个单一路径里，会根据需要添加斜杠。结果是经过简化的，所有的空字符串元素会被忽略。
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed, err:%v\n", err)
		return err
	}
	// error级别的日志，在原日志基础上添加 .err 即可
	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open err log file failed, err:%v\n", err)
		return err
	}
	// 日志文件都已经打开了
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return nil // 没有错误，就返回nil

}

// 定义函数，用来判断当前日志等级是否允许输出
func (f *FileLogger) enable(logLevel LogLevel) bool {
	return logLevel >= f.Lever
}

// 不同级别的日志中，输出内容的部分重复性太高，这里提取出来封装成方法
func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		// 字符串和参数拼接在一起才是完整的msg信息，我们这里直接使用原生的Sprintf来进行替换。
		msg := fmt.Sprintf(format, a...)
		now := time.Now().Format("2006-01-02 15:04:05")
		funcName, fileName, lineNo := getInfo(3)
		// getLogString 将日志级别数值转换成字符串
		// 这里是写入到文件，不再是控制台了
		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s\n", now, getLogString(lv), fileName, funcName, lineNo, msg)
		if lv >= ERROR {
			fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d] %s\n", now, getLogString(lv), fileName, funcName, lineNo, msg)
		}
	}
}

// Debug Debug级别日志,第一个是一个字符串，第二个是任意类型，数量不定的参数
func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}

// Trace Trace级别日志
func (f *FileLogger) Trace(format string, a ...interface{}) {
	f.log(TRACE, format, a...)
}

// Info Info级别日志
func (f *FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}

// Warning Warning级别日志
func (f *FileLogger) Warning(format string, a ...interface{}) {
	f.log(WARNING, format, a...)
}

// Error Error级别日志
func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}

// Fatal Fatal级别日志
func (f *FileLogger) Fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a...)
}
