package main

import "fmt"

func main() {
	// 输出正负号
	fmt.Printf("%+d\n", 10)
	fmt.Printf("%+d\n", -10)

	//数字证书前加空格，负数前加负号
	fmt.Printf("% d\n", 10)
	fmt.Printf("% d\n", -10)
	// 字符串 % x 输出会在打印的字节之间加空格
	fmt.Printf("% x\n", "哈哈哈")
	fmt.Printf("% X\n", "哈哈哈")

	// - 左对齐变成右对齐,这个案例如果没有 - 应该是左补空格，现在变成了右补空格
	fmt.Printf("%-11.4f\n", 10.02)

	// 使用0填充
	fmt.Printf("%011.4f\n", 10.02)

	// #
	fmt.Printf("%o\n", 10)
	fmt.Printf("%#o\n", 10) // 八进制前面添加0

	fmt.Printf("%x\n", 10)
	fmt.Printf("%#x\n", 10) // 十六进制添加0x

	var a = 10
	fmt.Printf("%p\n", &a)
	fmt.Printf("%#p\n", &a) // 指针去除前面的 0x

	fmt.Printf("%U\n", 97)
	fmt.Printf("%#U\n", 97)
}
