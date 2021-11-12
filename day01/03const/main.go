package main

import "fmt"

const (
	n1 = iota
	n2
)

const (
	n3 = iota // 这个iota不会受到上一个常量定义的影响，每次碰到const都会初始为0，所以这里也是从 0 开始
	n4
)

func main()  {
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)
}
