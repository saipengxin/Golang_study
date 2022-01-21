package main

import (
	"fmt"
	"github.com/shopify/sarama"
)

func main() {
	config := sarama.NewConfig()

	// Producer 生产者
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要 leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner //新选出⼀个 partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在 success channel返回

	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "web_log"
	msg.Value = sarama.StringEncoder("this is a log")

	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("producer closed,err", err)
		return
	}
	defer client.Close()

	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg fail,err:", err)
		return
	}
	fmt.Printf("pid:%v, offset:%v", pid, offset)
}
