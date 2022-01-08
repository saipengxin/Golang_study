package main

import (
	"fmt"
	"path"
)

func main() {
	// 	获取路径的最后一个元素
	var p1 = "/Users/saipengxin/Work/gopath/src/github.com/saipenxin/Golang_study/day04/Caller/main.go"
	fmt.Println(path.Base(p1))

	// 会将最后的 / 去除
	var p2 = "/a/b/c/"
	fmt.Println(path.Base(p2))

	var p3 = "/"
	fmt.Println(path.Base(p3))

	var p4 = ""
	fmt.Println(path.Base(p4))
}
