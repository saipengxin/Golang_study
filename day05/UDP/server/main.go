package main

import (
	"fmt"
	"net"
)

func main() {
	// 设置地址和端口
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	// 关闭要写在错误处理后面，如果出错了，那么返回的对象不一定实现了Close方法。
	defer listen.Close()

	for {
		// 直接读取数据就可以了
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("UDP读取失败, err:", err)
			continue
		}
		fmt.Printf("数据:%v 地址:%v 字节数:%v\n", string(data[:n]), addr, n)

		_, err = listen.WriteToUDP(data[:n], addr)
		if err != nil {
			fmt.Println("UDP消息发送失败, err:", err)
			continue
		}
	}
}
