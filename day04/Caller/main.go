package main

import (
	"fmt"
	"runtime"
)

func f1() {
	f2()
}

func f2() {
	f3()
}

func f3() {
	// 参数为2，表示上两层的调用位置，f3上一层是f2，也就是f2的调用位置
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		fmt.Println("错误")
		return
	}
	fmt.Println(file)
	fmt.Println(line)

	// 处理调用者标识符，并获取函数名
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName)
}

func main() {
	f1()
}
