package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/saipengxin/study/day08/go-micro/proto"
	"log"
)

type Hello struct{}

// Info 生成的方法在hello.pd.micro.go文件里，注意注释，这个方法会生成两个，一个是客户端的，一个是服务端的
// 我们这里写的是服务端的代码，所以要找到服务端的方法，然后按照他的参数和返回值定义
func (g *Hello) Info(ctx context.Context, req *proto.InfoRequest, rep *proto.InfoResponse) error {
	rep.Msg = "你好" + req.Username
	return nil
}

func main() {
	// 得到服务端实例
	service := micro.NewService(
		// 注册服务名称，要通过名称来调用
		micro.Name("hello"),
	)

	// 初始化
	service.Init()

	// 注册服务
	err := proto.RegisterHelloHandler(service.Server(), new(Hello))
	if err != nil {
		fmt.Println(err)
	}
	// 4.启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
