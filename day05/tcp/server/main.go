package main

import (
	"bufio"
	"fmt"
	"net"
)

// tcp 服务端

func main() {
	// 设置使用的协议，设置ip地址和监听的端口
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}

	// 循环处理 客户端请求，不加循环就成了一次性的了
	for {
		// Accept等待并返回下一个连接到该接口的连接,我们前面已经设置好了协议和监听的端口，接下来只要等待连接请求就可以了
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		// 连接成功，进行通信，但是我们不需要自己动手处理，交给小弟去处理
		// go中创建一个goroutine太方便了，所以我们将客户端的请求交给goroutine去处理，这里只负责监听就可以了
		// 一个客户端创建一个goroutine
		go process(conn)
	}

}

func process(conn net.Conn) {
	defer conn.Close() // 关闭连接
	//循环处理客户端的请求，客户端可能有多次请求，循环处理，不添加循环就成了一次性处理了
	for {
		reader := bufio.NewReader(conn) // 创建从请求中读取信息的对象，也可以使用conn自带的Read方法直接读取
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", recvStr)
		
		conn.Write([]byte(recvStr + "server端回传")) // 发送数据
	}
}
