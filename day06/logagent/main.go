package main

import (
	"fmt"
	"github.com/saipengxin/study/day06/logagent/conf"
	"github.com/saipengxin/study/day06/logagent/kafka"
	"github.com/saipengxin/study/day06/logagent/taillog"
	"gopkg.in/ini.v1"
	"time"
)

var (
	cfg = new(conf.AppConf)
)

// logAgent入口程序

func run() {
	// 读取日志
	for {
		select {
		case line := <-taillog.ReadChan():
			// 2. 发送到kafka
			kafka.SendToKafka(cfg.KafkaConf.Topic, line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}

// 入口函数
func main() {
	// 初始化配置文件
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Printf("load ini failed, err:%v\n", err)
		return
	}

	// 初始化kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address})
	if err != nil {
		fmt.Printf("init Kafka failed,err:%v\n", err)
		return
	}
	fmt.Println("init kafka success.")

	// 2. 打开日志文件准备收集日志
	err = taillog.Init(cfg.TaillogConf.FileName)
	if err != nil {
		fmt.Printf("Init taillog failed,err:%v\n", err)
		return
	}
	fmt.Println("init taillog success.")
	run()
}
