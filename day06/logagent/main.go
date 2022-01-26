package main

import (
	"fmt"
	"github.com/saipengxin/study/day06/logagent/conf"
	"github.com/saipengxin/study/day06/logagent/etcd"
	"github.com/saipengxin/study/day06/logagent/kafka"
	"github.com/saipengxin/study/day06/logagent/taillog"
	"github.com/saipengxin/study/day06/logagent/utils"
	"gopkg.in/ini.v1"
	"sync"
	"time"
)

var (
	cfg = new(conf.AppConf)
)

// logAgent入口程序

// 入口函数
func main() {
	// 初始化配置文件
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Printf("load ini failed, err:%v\n", err)
		return
	}

	// 初始化kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.ChanMaxSize)
	if err != nil {
		fmt.Printf("init Kafka failed,err:%v\n", err)
		return
	}
	fmt.Println("init kafka success.")

	// 初始化etcd
	err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.EtcdConf.Timeout)*time.Second)
	if err != nil {
		fmt.Printf("init etcd failed,err:%v\n", err)
		return
	}
	fmt.Println("init etcd success.")

	// 从etcd中获取日志收集项的配置信息
	// 为了实现每个logagent都拉取自己独有的配置，所以要以自己的IP地址作为区分
	ipStr, err := utils.GetOutboundIP()
	if err != nil {
		panic(err)
	}
	etcdConfKey := fmt.Sprintf(cfg.EtcdConf.Key, ipStr)
	logEntryConf, err := etcd.GetConf(etcdConfKey)
	if err != nil {
		fmt.Printf("etcd.GetConf failed,err:%v\n", err)
		return
	}
	fmt.Printf("get conf from etcd success, %v\n", logEntryConf)

	// 3. 收集日志发往Kafka
	taillog.Init(logEntryConf)

	// 监听新配置
	// 获取一个只写的通道，将新配置项写入这个通道中
	newConfChan := taillog.NewConfChan()
	var wg sync.WaitGroup
	wg.Add(1)
	go etcd.WatchConf(etcdConfKey, newConfChan) // 哨兵发现最新的配置信息会通知上面的那个通道
	wg.Wait()

}
