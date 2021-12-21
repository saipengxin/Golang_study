package main

import "fmt"

func f1(x func(int, int), m, n int) {
	x(m, n)
}

func main() {
	// 参数为函数类型的时候，我们可以直接传递匿名函数
	f1(func(x int, y int) {
		fmt.Println(x * y)
	}, 3, 4)
}
