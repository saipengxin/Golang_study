package main

import "fmt"

func main() {
	demo()
}

func demo() {
	fmt.Println("start")
	defer fmt.Println("哈哈哈")
	defer fmt.Println("嘿嘿嘿")
	defer fmt.Println("呵呵呵")
	fmt.Println("end")
}
