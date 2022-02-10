package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 参数
type ArithRequest struct {
	A, B int
}

// 用于注册
type Arith struct {
}

// 返回给客户端的结果
type ArithResponse struct {
	// 乘积
	Pro int
	// 商
	Quo int
	// 余数
	Rem int
}

func (a *Arith) Multiply(req ArithRequest, ret *ArithResponse) error {
	ret.Pro = req.A * req.B
	return nil
}

func (a *Arith) Divide(req ArithRequest, ret *ArithResponse) error {
	if req.B == 0 {
		return errors.New("除数不能为0")
	}
	ret.Quo = req.A / req.B
	ret.Rem = req.A / req.B
	return nil
}

func main() {
	// 注册一个服务
	rpc.Register(new(Arith))
	lis, err := net.Listen("tcp", ":8800")
	if err != nil {
		log.Fatal(err)
	}
	// 循环监听服务
	for {
		conn, err := lis.Accept()
		if err != nil {
			continue
		}
		go func(conn net.Conn) {
			fmt.Println("新请求")
			jsonrpc.ServeConn(conn)
		}(conn)
	}
}
