package main

import (
	"fmt"
	"io"
	"os"
)

//
// CopyFile
//  @Description: 文件复制
//  @param old_file 要复制的文件
//  @param new_file 复制出来的新文件
//  @return written 拷贝的字节数
//  @return err 遇到的错误
//
func CopyFile(old_file, new_file string) (written int64, err error) {
	// 以只读的方式打开要复制的文件
	reader, err := os.Open(old_file)
	if err != nil {
		fmt.Printf("打开%s文件出错\n", old_file)
		return
	}
	// 关闭文件
	defer reader.Close()

	// 以 覆盖，只写的方式打开要复制到的文件，如果没有就创建
	writer, err := os.OpenFile(new_file, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Printf("打开%s文件出错\n", new_file)
		return
	}
	defer writer.Close()

	// 复制内容，将reader的内容复制到writer
	return io.Copy(writer, reader)
}

func main() {
	written, err := CopyFile("readme.txt", "new_readme.txt")
	if err != nil {
		fmt.Println("出错误了：", err)
		return
	}
	fmt.Printf("拷贝了%d字节", written)
}
