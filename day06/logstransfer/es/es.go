package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"strings"
	"time"
)

type LogData struct {
	Topic string `json:"topic"`
	Data  string `json:"data"`
}

var (
	client *elastic.Client
	ch     chan *LogData
)

// Init Es初始化
func Init(address string, chanSize, nums int) (err error) {
	if !strings.HasPrefix(address, "http://") {
		address = "http://" + address
	}
	client, err = elastic.NewClient(elastic.SetURL(address))
	if err != nil {
		// Handle error
		return err
	}
	fmt.Println("connect to es success")
	ch = make(chan *LogData, chanSize)
	for i := 0; i < nums; i++ {
		go sendToES()
	}
	return
}

// SendToES 发送数据到ES
func SendToESChan(msg *LogData) {
	ch <- msg
}

func sendToES() {
	// 链式操作
	for {
		select {
		case msg := <-ch:
			put1, err := client.Index().Index(msg.Topic).BodyJson(msg).Do(context.Background())
			if err != nil {
				// Handle error
				fmt.Println(err)
				continue
			}
			fmt.Printf("Indexed student %v to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
		default:
			time.Sleep(time.Second)
		}
	}
}
