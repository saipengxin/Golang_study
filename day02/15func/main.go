package main

import "fmt"

func main() {
	var str = "上海自来水来自海上"

	// 先将字符串转为切片,否则无法使用for循环遍历出汉字来
	var str_slice = []rune(str)

	// 遍历切片
	for i := 0; i < len(str_slice)/2; i++ {
		if str_slice[i] != str_slice[len(str_slice)-1-i] {
			fmt.Println("不是回文")
		} else {
			fmt.Println("是回文")
		}
	}
}
