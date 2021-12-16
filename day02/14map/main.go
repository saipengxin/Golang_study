package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = "how do you do"

	// 将字符串使用 空格 分割
	var str_slice = strings.Split(str, " ")
	fmt.Printf("%T", str_slice)

	// 定义一个map类型的变量，用于保存最后的结果
	var res = make(map[string]int, 10)

	// 遍历切片，统计单词数量
	for _, v := range str_slice {
		// 根据 go语言的特性，如果初始化的时候没有指定值会初始化对应类型的零值，所以这里初始化了int的零值，也就是0.
		// 所以我们可以直接哪来计算，无需重新赋值0
		res[v] += 1
	}

	fmt.Println(res)
}
