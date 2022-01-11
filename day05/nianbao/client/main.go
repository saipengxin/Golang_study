package main

import (
	"fmt"
	"github.com/saipengxin/study/day05/nianbao/protocol"
	"net"
	"strconv"
)

// socket_stick/client/main.go

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello, Hello. How are you?` + strconv.Itoa(i)
		// 客户端发送数据前使用Encode函数进行编码
		data, err := protocol.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed, err:", err)
			return
		}
		conn.Write(data)

	}
}
