package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var a = new(person)
	// 因为go允许直接使用结构体指针配合 . 语法来获取成员变量
	a.name = "saipx"
	a.age = 15

	fmt.Printf("%#v\n", a)
	fmt.Printf("%T\n", a)

	var b = new(person)
	// 给结构体指针初始化，前面要记得带上&,因为b是指针类型的变量，所以等号右边也要是指针类型的数据，所以要使用&
	b = &person{
		name: "saipx",
		age:  15,
	}
	fmt.Printf("%#v\n", b)
	fmt.Printf("%T\n", b)
}
