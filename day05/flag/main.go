package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var name string
	var age int
	var delay time.Duration
	flag.StringVar(&name, "name", "saipx", "请输入姓名")           // 第一个参数传递的是变量的指针
	flag.IntVar(&age, "age", 18, "请输入年龄")                     // 第一个参数传递的是变量的指针
	flag.DurationVar(&delay, "delay", time.Second, "请输入时间间隔") // 第一个参数传递的是变量的指针

	flag.Parse() // 解析flag参数，这一步是必须的
	fmt.Println("name=", name)
	fmt.Println("age=", age)
	fmt.Println("delay=", delay)

	fmt.Println(flag.Args())  // 返回命令行参数后的其他参数，[]string类型
	fmt.Println(flag.NArg())  // 返回命令行参数后的其他参数个数
	fmt.Println(flag.NFlag()) // 返回使用的命令行参数个数
}
