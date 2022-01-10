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
		data, err := protocol.Encode(msg)
		fmt.Println(data, 11)
		if err != nil {
			fmt.Println("encode msg failed, err:", err)
			return
		}
		conn.Write(data)

	}
}
