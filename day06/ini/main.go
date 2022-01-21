package main

import (
	"fmt"
	"gopkg.in/ini.v1"
)

// AppConf 包含KafkaConf和TaillogConf两个匿名结构体
type AppConf struct {
	KafkaConf   `ini:"kafka"`
	TaillogConf `ini:"taillog"`
}
type KafkaConf struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

type TaillogConf struct {
	FileName string `ini:"filename"`
}

func main() {
	// 获取结构体指针
	cfg := new(AppConf)
	// 将配置文件的信息映射给结构体
	err := ini.MapTo(cfg, "./config.ini")
	if err != nil {
		fmt.Printf("load ini failed, err:%v\n", err)
		return
	}
	fmt.Printf("%#v\n", cfg)
}
