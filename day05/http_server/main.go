package main

import (
	"fmt"
	"net/http"
)

func main() {

	// 给 /user/info 添加处理程序，使用匿名函数的方式
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		// 给客户端响应数据
		answer := "ok"
		w.Write([]byte(answer))

		// 解析get请求的参数
		data := r.URL.Query()
		fmt.Println(data.Get("name"))
		fmt.Println(data.Get("age"))

		fmt.Println(r.URL)    // 请求地址
		fmt.Println(r.Body)   // body信息，post会使用body传递参数，get方式的参数在URL上
		fmt.Println(r.Method) // 请求方式

	})

	// 启动一个http服务端
	err := http.ListenAndServe("127.0.0.1:9090", nil)
	if err != nil {
		fmt.Println("报错：", err)
	}

}
