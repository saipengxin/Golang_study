package main

import "fmt"

func f5() func(int, string) int {
	return f6
}

func f6(x int, y string) int {
	return 3
}

func main() {
	a := f5()
	fmt.Printf("%T\n", a)

	b := a(1, "哈")
	fmt.Println(b)
}
