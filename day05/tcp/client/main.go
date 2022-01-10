package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// tcp 客户端
func main() {
	//Dial在指定的网络上连接指定的地址。用于主动请求，Listen是用于监听
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("err :", err)
		return
	}
	defer conn.Close() // 关闭连接

	inputReader := bufio.NewReader(os.Stdin) // 创建从标准输入的对象
	for {
		fmt.Print("请输入:")
		input, _ := inputReader.ReadString('\n')  // 读取用户输入,换行结束
		inputInfo := strings.Trim(input, "\r\n")  // 去除内容的换行
		if strings.ToUpper(inputInfo) == "exit" { // 如果输入exit就退出程序
			return
		}
		_, err = conn.Write([]byte(inputInfo)) // 发送数据，使用Write方法，向对象中写入数据
		if err != nil {
			return
		}

		buf := [512]byte{}
		n, err := conn.Read(buf[:]) // Read从连接中读取数据
		if err != nil {
			fmt.Println("recv failed, err:", err)
			return
		}
		fmt.Println("收到server端回传数据", string(buf[:n]))
	}
}
