package main

import (
	"context"
	"fmt"
	"github.com/saipengxin/study/day08/gRPC/proto"
	"google.golang.org/grpc"
)

func main() {
	// 连接
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("连接异常:%s\n", err)
	}
	defer conn.Close()

	// 实例化grpc客户端
	client := proto.NewUserInfoServiceClient(conn)
	// 组装请求参数
	req := new(proto.UserRequest)
	req.Name = "zs"
	// 调用接口
	response, err := client.GetUserInfo(context.Background(), req)
	if err != nil {
		fmt.Println("响应异常 %s\n", err)
		return
	}
	fmt.Printf("响应结果：%v\n", response)
}
