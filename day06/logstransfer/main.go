package main

import (
	"fmt"
	"github.com/saipengxin/study/day06/logstransfer/config"
	"github.com/saipengxin/study/day06/logstransfer/es"
	"github.com/saipengxin/study/day06/logstransfer/kafka"
	"gopkg.in/ini.v1"
)

func main() {
	// 加载配置文件
	var cfg = new(config.LogTransferCfg)
	err := ini.MapTo(cfg, "./config/config.ini")
	if err != nil {
		fmt.Println("ini load fail，err:", err)
		return
	}

	// 1. 初始化ES
	// 1.1 初始化一个ES连接的client
	// 1.2 对外提供一个往ES写入数据 的一个函数
	err = es.Init(cfg.ESCfg.Address, cfg.ESCfg.ChanSize, cfg.ESCfg.Nums)
	if err != nil {
		fmt.Printf("init ES client failed,err:%v\n", err)
		return
	}
	fmt.Println("init es success.")

	// 2. 初始化kafka
	// 2.1 连接kafka, 创建分区的消费者
	// 2.2 每个分区的消费者分别取出数据 通过SendToES()将数据发往ES
	err = kafka.Init([]string{cfg.KafkaCfg.Address}, cfg.KafkaCfg.Topic)
	if err != nil {
		fmt.Printf("init kafka consumer failed,err:%v\n", err)
		return
	}
	select {}
}
