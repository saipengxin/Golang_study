package main

import (
	"context"
	"fmt"
	"github.com/saipengxin/study/day08/gRPC/proto"
	"google.golang.org/grpc"
	"net"
)

// 定义空接口
type UserInfoService struct{}

var u = UserInfoService{}

// 实现在 proto 文件中定义的方法，我们这里是在服务端实现，要使用服务端的接口中定义的方法，参数要同步
func (u UserInfoService) GetUserInfo(ctx context.Context, req *proto.UserRequest) (resp *proto.UserResponse, err error) {
	name := req.Name
	if name == "zs" {
		resp = &proto.UserResponse{
			Id:    1,
			Name:  name,
			Age:   22,
			Hobby: []string{"Sing", "Sun"},
		}
	}
	return
}

func main() {
	// 地址
	addr := "127.0.0.1:8080"
	// 监听
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("监听异常:%s\n", err)
	}
	fmt.Printf("监听端口：%s\n", addr)

	// 实例化GRPC
	s := grpc.NewServer()
	// 在grpc上注册微服务
	proto.RegisterUserInfoServiceServer(s, &u)
	// 启动服务端
	s.Serve(listener)
}
