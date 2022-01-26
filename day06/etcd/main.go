package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"}, // 连接的etcd集群地址，这里为单机的故一个地址
		DialTimeout: 5 * time.Second,            // 超时时长
	})

	if err != nil {
		fmt.Printf("connect to etcd faild,err,%v\n", err)
		return
	}
	fmt.Println("connect success")
	defer cli.Close()

	// put
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	log := `[{"path":"D:/text/nginx.log","topic":"web_log"}]`
	_, err = cli.Put(ctx, "/logagent/collect_config", log)
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}

	// get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "name", clientv3.WithPrefix())
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	//Kvs是一个切片，遍历
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}

}
