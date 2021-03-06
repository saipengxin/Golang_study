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
	maxFileSize int64        // 文件切割使用
	logChan     chan *logMsg // 保存日志的通道
}

// 通道要保存的日志数据
type logMsg struct {
	level     LogLevel
	msg       string
	funcName  string
	fileName  string
	timestamp string
	line      int
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
		logChan:     make(chan *logMsg, 10000), // 初始化10000个缓冲空间
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
	// 启动一个goroutine
	go f.writeLogBackground()
	return nil // 没有错误，就返回nil

}

// 定义方法，用来判断当前日志等级是否允许输出
func (f *FileLogger) enable(logLevel LogLevel) bool {
	return logLevel >= f.Lever
}

// 定义方法，判断是否需要切割
func (f *FileLogger) checkSize(file *os.File) bool {
	// 因为我们有两个日志文件，普通日志文件和error级别的日志文件，所以我们要传递文件句柄，根据传递的文件句柄来判断不同的文件，不能直接从接收者f中获取。
	// 因为我们无法确定取哪个
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("获取文件信息失败")
		return false
	}
	return fileInfo.Size() >= f.maxFileSize
}

// 定义文件切割方法
func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	// 还是因为存在不同的日志文件，所以我要传递文件句柄，根据文件句柄来切割不同的文件
	// 1.根据当前文件句柄，生成备份文件名称 xx.log ==> xx.log.bak20220106172100
	nowStr := time.Now().Format("20060102150405")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("获取文件信息失败")
		return nil, err
	}
	logName := path.Join(f.filePath, fileInfo.Name())      // 文件名要实时获取，不同的日志文件名不同，这里获取日志文件的全路径+ 文件名
	newLogName := fmt.Sprintf("%s.bak%s", logName, nowStr) // 备份的文件路径 + 文件名

	// 2.关闭当前日志
	file.Close()
	// 3、将日志备份
	os.Rename(logName, newLogName)
	// 4. 打开一个新的日志文件
	// 我们的logName上一步已经成功改名了，这里直接打开新的logName就可以了
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open new log file failed, err:%v\n", err)
		return nil, err
	}
	// 5. 将打开的新日志文件返回
	return fileObj, nil
}

// 将日志写入文件，提取出来封装一下，启动goroutine后调用这个写入文件
func (f *FileLogger) writeLogBackground() {
	for {
		// 判断是否需要切割日志
		if f.checkSize(f.fileObj) {
			newFile, err := f.splitFile(f.fileObj)
			if err != nil {
				return
			}
			f.fileObj = newFile
		}

		// 从通道中读取日志内容
		select {
		case logTmp := <-f.logChan:
			// 将日志信息拼接出来
			logInfo := fmt.Sprintf("[%s] [%s] [%s:%s:%d] %s\n", logTmp.timestamp, getLogString(logTmp.level), logTmp.fileName, logTmp.funcName, logTmp.line, logTmp.msg)

			// getLogString 将日志级别数值转换成字符串
			// 这里是写入到文件，不再是控制台了
			fmt.Fprintf(f.fileObj, logInfo)
			if logTmp.level >= ERROR {
				// 判断日志是否需要切割
				if f.checkSize(f.errFileObj) {
					newFile, err := f.splitFile(f.errFileObj) // 日志文件
					if err != nil {
						return
					}
					f.errFileObj = newFile
				}
				fmt.Fprintf(f.errFileObj, logInfo)
			}
		default:
			// 取不到日志先休息500毫秒,sleep的时候是会让出cpu的。
			time.Sleep(time.Millisecond * 500)

		}
	}
}

// 不同级别的日志中，输出内容的部分重复性太高，这里提取出来封装成方法
func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		// 字符串和参数拼接在一起才是完整的msg信息，我们这里直接使用原生的Sprintf来进行替换。
		msg := fmt.Sprintf(format, a...)
		now := time.Now().Format("2006-01-02 15:04:05")
		funcName, fileName, lineNo := getInfo(3)

		// 将日志写入通道
		// 构造一个logMsg的对象
		l := &logMsg{
			level:     lv,
			msg:       msg,
			funcName:  funcName,
			fileName:  fileName,
			timestamp: now,
			line:      lineNo,
		}
		// 写入通道
		select {
		case f.logChan <- l:
		default:
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
