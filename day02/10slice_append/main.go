package main

import "fmt"

func main()  {
	var a = make([]int,3,6)
	fmt.Println(a)

	b := append(a,4,5,6,7)
	fmt.Println(b)

	b[0] = 100
	fmt.Printf("%p,%v\n",a,a)
	fmt.Printf("%p,%v\n",b,b)
}
