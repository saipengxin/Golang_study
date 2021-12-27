package main

import "fmt"

type person struct {
	name string
	city string
	age  int
}

// 构造函数
func NewPerson(name, city string, age int) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

//方法
// 前面这个括号是指定的接收者，一般用接收者类型的第一个字母命名
func (p person) Eating() {
	fmt.Printf("%s正在吃饭", p.name)
}

func main() {
	p := NewPerson("saipx", "北京", 18)
	fmt.Println(p)
	// 调用方法，因为指定了接收者为person类型，所以只有person类型的变量才能调用
	p.Eating()

	// 不能像普通方法一样直接调用
	// Eating()

}
