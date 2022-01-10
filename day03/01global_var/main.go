package main

import "fmt"

// 全局作用域

func main() {
	if i := 0; i == 3 {
		fmt.Println(i)
	}
	//fmt.Println()
}
