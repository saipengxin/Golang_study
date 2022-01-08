package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("请输入内容：")
	// 打开一个读对象，读的来源是标准输入，也就是控制台，os.Stdin
	reader := bufio.NewReader(os.Stdin)
	// 读取控制台输入的内容，换行结束
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("报错了")
		return
	}
	fmt.Println("您输入的是：", text)
}
