package main

import (
	"fmt"
	"io"
	"os"
)

// 从 文件的 5字节 后的位置进行内容插入
func main() {
	// 打开原文件 只读方式
	file, err := os.OpenFile("demo.txt", os.O_RDONLY, 0777)
	if err != nil {
		fmt.Println("文件打开失败")
		return
	}
	defer file.Close()

	// 建立临时文件
	filetmp, err := os.Create("demo.tmp")
	if err != nil {
		fmt.Println("文件创建失败")
		return
	}
	defer filetmp.Close()

	// 从原文件读取指定字节内容写入临时文件
	var b [5]byte
	n, err := file.Read(b[:])
	if err != nil {
		fmt.Println("读取失败1")
		return
	}
	filetmp.Write(b[:n])

	// 将要插入的内容写入临时文件
	var s = "sai"
	filetmp.Write([]byte(s))

	// 将原文件的后续内容写入临时文件
	var x [1024]byte
	for {
		// 他会接着上次读取的位置继续读
		n, err := file.Read(x[:])
		if err == io.EOF {
			filetmp.Write(x[:n])
			fmt.Println("文件读取完成")
			break
		}
		if err != nil {
			fmt.Println("读取失败2", err)
			break
		}

		filetmp.Write(x[:n])
	}

	// 使用临时文件替换原文件
	err = os.Rename("demo.tmp", "demo.txt")
	if err != nil {
		fmt.Println("失败", err)
	}

}
