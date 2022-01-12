package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type person struct {
	Name string
	Aga  int
}

func main() {
	var apiUrl = "http://127.0.0.1:9090"

	// 方式1、http.Post
	// 数据类型
	//contentType := "application/json"
	//// 传输的数据
	//data := `{"name":"saipx","age":18}` // json格式字符串
	//resp, err := http.Post(apiUrl, contentType, strings.NewReader(data))

	// 方式2：http.NewRequest
	p := person{"saipx", 18}
	// 转为 json 串
	data, err := json.Marshal(p)
	reader := bytes.NewReader(data)
	request, err := http.NewRequest("POST", apiUrl, reader)
	request.Header.Set("Content-Type", "application/json") // 设置传送数据类型为json
	resp, err := http.DefaultClient.Do(request)

	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}
