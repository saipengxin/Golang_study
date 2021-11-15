package main

import (
	"fmt"
	"strings"
)

func main() {
	var name string = "赛鹏新"
	var str string = "saipengxin"
	var path string = "D:\\saipx\\golang\\day01"

	// 1、len() 获取字符串长度
	fmt.Printf("name长度为 %d 个字节\n",len(name))
	fmt.Printf("str长度为 %d 个字节\n",len(str))
	fmt.Println("============================")

	// 2、 字符串拼接
	// (1)、 +
	var str1 = "我是" + str
	fmt.Println(str1)
	// (2)、fmt.Sprintf
	var str2 = fmt.Sprintf("我是%s",str)
	fmt.Println(str2)
	fmt.Println("============================")

	// 3、字符串分割
	fmt.Println(strings.Split(path,"\\"))
	fmt.Println("============================")

	// 4、字符串是否包含
	fmt.Println(strings.Contains(str,"s"))
	fmt.Println(strings.Contains(str,"xx"))
	fmt.Println("============================")

	// 5、前缀 / 后缀判断
	fmt.Println(strings.HasPrefix(str,"s")) // 前缀
	fmt.Println(strings.HasSuffix(str,"n")) // 后缀
	fmt.Println("============================")

	// 6、子串出现的位置
	fmt.Println(strings.Index(str,"i"))
	fmt.Println(strings.LastIndex(str,"i"))
	fmt.Println("============================")

	// 字符串连接
	fmt.Println(strings.Join(strings.Split(path,"\\"),"+")) // 将使用 \ 分割的内容 使用 + 拼接起来

}
