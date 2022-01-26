package taillog

import (
	"fmt"
	"github.com/saipengxin/study/day06/logagent/etcd"
	"time"
)

// 定义一个全局的管理者
var tskMgr *tailLogMgr

// 任务管理对象，这里存储的是每一个任务的配置项
type tailLogMgr struct {
	logEntry    []*etcd.LogEntry
	tskMap      map[string]*TailTask
	newConfChan chan []*etcd.LogEntry // 新配置通知的通道
}

func Init(logEntryConf []*etcd.LogEntry) {
	tskMgr = &tailLogMgr{
		logEntry:    logEntryConf, // 把当前的日志收集项配置信息保存起来
		tskMap:      make(map[string]*TailTask, 16),
		newConfChan: make(chan []*etcd.LogEntry),
	}

	for _, logEntry := range logEntryConf {
		//conf: *etcd.LogEntry
		//logEntry.Path： 要收集的日志文件的路径
		tailObj := NewTailTask(logEntry.Path, logEntry.Topic)
		mk := fmt.Sprintf("%s_%s", logEntry.Path, logEntry.Topic)
		tskMgr.tskMap[mk] = tailObj
	}

	go tskMgr.run()
}

func (t *tailLogMgr) run() {
	for {
		select {
		case newConf := <-t.newConfChan:
			for _, conf := range newConf {
				mk := fmt.Sprintf("%s_%s", conf.Path, conf.Topic)
				_, ok := t.tskMap[mk]
				if ok {
					// 原来就有，不需要操作
					continue
				} else {
					// 新增的
					tailObj := NewTailTask(conf.Path, conf.Topic)
					t.tskMap[mk] = tailObj
				}
			}
			// 找出原来t.logEntry有，但是newConf中没有的，要删掉
			for _, c1 := range t.logEntry { // 从原来的配置中依次拿出配置项
				isDelete := true
				for _, c2 := range newConf { // 去新的配置中逐一进行比较
					if c2.Path == c1.Path && c2.Topic == c1.Topic {
						isDelete = false
						continue
					}
				}
				if isDelete {
					// 把c1对应的这个tailObj给停掉
					mk := fmt.Sprintf("%s_%s", c1.Path, c1.Topic)
					//t.tskMap[mk] ==> tailObj
					t.tskMap[mk].cancelFunc()
				}
			}
		default:
			time.Sleep(time.Second)
		}
	}
}

// NewConfChan 一个函数，向外暴露tskMgr的newConfChan,只写，可以将新配置写入通道中
func NewConfChan() chan<- []*etcd.LogEntry {
	return tskMgr.newConfChan
}
