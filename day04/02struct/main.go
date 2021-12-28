package main

import "fmt"

//Animal 动物
type Animal struct {
	name string
}

// Animal指针类型的接收者，这个方法只有Animal类型的变量才能调用
func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

//Dog 狗
type Dog struct {
	Feet    int8
	*Animal //通过嵌套匿名字段实现继承，此时Dog不光继承了Animal结构体中的参数，还继承了他的方法
}

func (d *Dog) wang() {
	// 这里先在Dog中找name，没有找到会去嵌套的匿名字段中找name，也就是Animal
	fmt.Printf("%s会汪汪汪~\n", d.name)
}

func main() {
	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{ //注意嵌套的是结构体指针
			name: "乐乐",
		},
	}
	d1.wang() //乐乐会汪汪汪~
	d1.move() //乐乐会动！
}
