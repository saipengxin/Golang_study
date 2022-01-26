package taillog

import (
	"context"
	"fmt"
	"github.com/hpcloud/tail"
	"github.com/saipengxin/study/day06/logagent/kafka"
)

// 专门从日志文件收集日志的模块

// TailTask 一个日志收集的任务
type TailTask struct {
	path     string
	topic    string
	instance *tail.Tail
	// 为了能实现退出t.run()
	ctx        context.Context
	cancelFunc context.CancelFunc
}

func NewTailTask(path, topic string) (tailObj *TailTask) {
	ctx, cancel := context.WithCancel(context.Background())
	tailObj = &TailTask{
		path:       path,
		topic:      topic,
		ctx:        ctx,
		cancelFunc: cancel,
	}
	tailObj.init() // 根据路径去打开对应的日志
	return
}

// Init 初始化文件监听
func (t TailTask) init() {
	config := tail.Config{
		ReOpen:    false,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}
	var err error
	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Println("tail file err : ", err)
	}
	go t.run() // 直接去采集日志发送到kafka
}

func (t *TailTask) run() {
	for {
		select {
		case <-t.ctx.Done():
			fmt.Printf("tail task:%s_%s 结束了...\n", t.path, t.topic)
			return
		case line := <-t.instance.Lines: // 从tailObj的通道中一行一行的读取日志数据
			// 发往Kafka
			// 这里有一个问题，那就是函数调用函数，要等到上一个函数返回，才能执行下一次的循环，
			// 所以我们这里决定，不是直接将日志发送到kafka中，而是发送到一个通道中，然后我们在kafka中写一个从通道中读取日志的方法
			// 发送到通道中是很快的，就避免了直接发送到kafka可能会造成的延时
			kafka.SendToKafka(t.topic, line.Text)
			// kafka那个包中有单独的goroutine去取日志数据发到kafka
		}
	}
}
