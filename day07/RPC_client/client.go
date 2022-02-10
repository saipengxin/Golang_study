package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

// 参数
type ArithRequest struct {
	A, B int
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

func main() {
	// 连接远程的rpc服务
	conn, err := jsonrpc.Dial("tcp", ":8800")
	if err != nil {
		log.Fatal(err)
	}

	var res ArithResponse
	err = conn.Call("Arith.Multiply", ArithRequest{50, 10}, &res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("乘积:", res.Pro)

	err = conn.Call("Arith.Divide", ArithRequest{50, 10}, &res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("商:%d,余数%d\n", res.Quo, res.Rem)

}
