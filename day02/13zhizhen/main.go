package main

import "fmt"

func main()  {
	var a []int
	fmt.Println(a == nil)
	a[0] = 1
	fmt.Println(a)
}
