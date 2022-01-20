package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "./xx.log"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("打开失败", err)
		return
	}

	file.Seek(0, 2)

	var b [4]byte
	n, err := file.Read(b[:])
	if err != nil {
		fmt.Println("读取失败", err)
		return
	}
	fmt.Println(string(b[:n]))

}
