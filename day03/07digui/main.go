package main

import "fmt"

func f(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}
	return f(n-1) + f(n-2)
}

func main() {
	res := f(6)
	fmt.Println(res)
}
