package main

import "fmt"

func main() {
	demo("123412")
}

func demo(a interface{}) {
	switch a.(type) {
	case string:
		fmt.Println("string类型")
	case int64:
		fmt.Println("int64类型")
	case float32:
		fmt.Println("float32类型")
	default:
		fmt.Println("其他类型")
	}

	// 如我们希望使用这个值，可以使用变量接受一下
	switch t := a.(type) {
	case string:
		fmt.Println("string类型", t)
	case int64:
		fmt.Println("int64类型", t)
	case float32:
		fmt.Println("float32类型", t)
	default:
		fmt.Println("其他类型", t)
	}
}
