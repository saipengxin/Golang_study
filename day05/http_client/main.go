package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	//resp, err := http.Get("127.0.0.1:9090?name=saipx&age=18")

	// 上面的方式就是携带参数的get请求，但是考虑到参数中可能会有各种特殊字符，我们一般会使用下面的方法来实现
	var apiUrl = "http://127.0.0.1:9090"

	// Values 是url包中自定义的一个类型，底层类型是 map[string][]string，用来存储要传输的数据
	data := url.Values{}
	data.Set("name", "saipx")
	data.Set("age", "20")

	// ParseRequestURI函数解析apiUrl为一个URL结构体
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Printf("parse url requestUrl failed, err:%v\n", err)
		return
	}

	// 数据编码并赋值
	// RawQuery 编码后的查询字符串，没有'?'
	u.RawQuery = data.Encode()

	// 接下来两种方式可以发送请求
	// 1.http.Get()
	//resp, err := http.Get(u.String())

	// 2、http.NewRequest
	req, err := http.NewRequest("GET", u.String(), nil)
	resp, err := http.DefaultClient.Do(req) // 发送请求
	if err != nil {
		fmt.Printf("req failed, err:%v\n", err)
		return
	}

	defer resp.Body.Close()
	// 从resp中将返回数据读取出来
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}
