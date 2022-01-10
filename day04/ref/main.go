package main

import (
	"fmt"
	r "github.com/saipengxin/study/day04/reflect_demo"
)

func main() {
	r.Check_action(10)

	type book struct{}
	r.Check_action(book{})

	r.Check_action(person{})
}

type person struct{}

func (p person) Action() {
	fmt.Println("action方法")
}

func (p person) Before_action() {
	fmt.Println("前置方法")
}
