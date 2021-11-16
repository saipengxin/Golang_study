package main

import "fmt"

func main()  {
	var a = "哈哈哈"
	fmt.Println(a)
	fmt.Println(&a)
	a = "嘿嘿嘿"
	fmt.Println(a)
	fmt.Println(&a)
}
