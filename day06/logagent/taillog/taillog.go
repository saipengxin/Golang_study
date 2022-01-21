package taillog

import (
	"fmt"
	"github.com/hpcloud/tail"
)

// 专门从日志文件收集日志的模块

var (
	tailObj *tail.Tail  // 监听文件的对象
	LogChan chan string // 将信息发送到通道中，另一边使用通道进行接收
)

// Init 初始化文件监听
func Init(fileName string) (err error) {
	config := tail.Config{
		ReOpen:    false,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}
	tailObj, err = tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail file err : ", err)
		return
	}
	return
}

// ReadChan 返回一个只读的通道类型，读取数据
func ReadChan() <-chan *tail.Line {
	return tailObj.Lines
}
