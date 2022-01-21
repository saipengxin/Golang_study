package kafka

import (
	"fmt"
	"github.com/shopify/sarama"
)

// 定义一个全局的kafka连接对象
var (
	client sarama.SyncProducer // 声明一个全局的连接kafka的生产者client =
)

// Init 注意这个大写的Init和我们前面学的init函数是不同的，这就是一个普通的自定义函数
func Init(addrs []string) (err error) {
	config := sarama.NewConfig()

	// Producer 生产者
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要 leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner //新选出⼀个 partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在 success channel返回

	// 连接kafka
	client, err = sarama.NewSyncProducer(addrs, config)
	if err != nil {
		fmt.Println("producer closed,err", err)
		return
	}
	return
}

// SendToKafka 向kafka中写入数据
// topic : 消息主题
// data  : 消息内容
func SendToKafka(topic, data string) {
	// 构造⼀个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)

	// 发送到kafka
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
