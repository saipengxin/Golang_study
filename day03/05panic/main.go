package main

import "fmt"

func f1() {
	fmt.Println("这是f1")
}

func f2() {
	defer func() {
		res := recover()
		if res != nil {
			// 返回值不是nil，证明触发了panic，并且捕获到了panic的输出信息
			fmt.Println("触发了panic，报错信息为：", res)
			// 我们可以在这个判断里写一些补救程序的操作，或者是关闭已经打开的资源，防止占用，又或者直接返回一个友好的提示信息
			fmt.Println("我们写了一些补救操作，程序继续执行了。")
		}
	}()
	panic("发生了错误，宕机！！！")
	fmt.Println("这是f2")
}

func f3() {
	fmt.Println("这里是f3")
}

func main() {
	f1()
	f2()
	f3()
}
