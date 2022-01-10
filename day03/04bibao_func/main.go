package main

import "fmt"

func f1() func(int) {
	var x = 10
	return func(y int) {
		x += y
		fmt.Printf("%p\n", x)
	}
}

func f2(y int) {
	var x = 10
	x = x + y
	fmt.Printf("%T\n", x)
}

func main() {
	var f = f1()
	f(10)
	f(20)
	fmt.Println("======================")
	var ff = f1()
	ff(10)
	ff(20)
	fmt.Println("======================")
	var fff = f2
	fff(10)
	fff(20)
}
