package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		defer request.Body.Close()

		// 打印post数据
		// 1. 请求类型是application/x-www-form-urlencoded时解析form数据
		request.ParseForm()
		fmt.Println(request.PostForm)
		fmt.Println(request.PostForm.Get("name"), request.PostForm.Get("age"))

		// 2. 请求类型是application/json时从r.Body读取数据
		b, err := ioutil.ReadAll(request.Body)
		if err != nil {
			fmt.Printf("read request.Body failed, err:%v\n", err)
			return
		}
		fmt.Println(string(b))

		// 回复消息
		writer.Write([]byte("请求成功"))
	})

	http.ListenAndServe("127.0.0.1:9090", nil)
}
