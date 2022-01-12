package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// 连接
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("连接服务端失败，err:", err)
		return
	}
	defer socket.Close()

	// 实例化 标准输入 对象
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("请输入：")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("获取输入错误", err)
			return
		}
		_, err = socket.Write([]byte(line))
		if err != nil {
			fmt.Println("发送数据失败，err:", err)
			return
		}

		data := make([]byte, 4096)
		n, remoteAddr, err := socket.ReadFromUDP(data) // 接收数据
		if err != nil {
			fmt.Println("接收数据失败，err:", err)
			return
		}
		fmt.Printf("接收数据:%v 地址:%v 字节数:%v\n\n", string(data[:n]), remoteAddr, n)
	}

}
