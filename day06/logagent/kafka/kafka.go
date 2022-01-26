package kafka

import (
	"fmt"
	"github.com/shopify/sarama"
	"time"
)

type logData struct {
	topic string
	data  string
}

// 定义一个全局的kafka连接对象
var (
	client      sarama.SyncProducer // 声明一个全局的连接kafka的生产者client
	logDataChan chan *logData
)

// Init 注意这个大写的Init和我们前面学的init函数是不同的，这就是一个普通的自定义函数
func Init(addrs []string, maxSize int) (err error) {
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
	// 初始化lodDataChan
	logDataChan = make(chan *logData, maxSize)
	// 开启后台的goroutine从通道中取取数据发往kafka
	go sendToKafka()

	return
}

// SendToKafka 给外部暴露的一个函数，该函数只把日志数据发送到一个内部的channel中
func SendToKafka(topic, data string) {
	msg := &logData{
		topic: topic,
		data:  data,
	}
	logDataChan <- msg
}

// 真正往kafka发送日志的函数
func sendToKafka() {
	for {
		select {
		case ld := <-logDataChan:
			// 构造⼀个消息
			msg := &sarama.ProducerMessage{}
			msg.Topic = ld.topic
			msg.Value = sarama.StringEncoder(ld.data)
			// 发送到kafka
			pid, offset, err := client.SendMessage(msg)
			if err != nil {
				fmt.Println("send msg failed, err:", err)
				return
			}
			fmt.Printf("pid:%v offset:%v\n", pid, offset)
		default:
			time.Sleep(time.Millisecond * 50)
		}

	}

}
