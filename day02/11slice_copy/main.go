package main

import "fmt"

func main()  {
	var a = [][]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}

	var b = make([][]int,len(a))
	for i := range a {
		b[i] = make([]int,len(a[i]))
		copy(b[i],a[i])
	}
	fmt.Println(a)
	fmt.Println(b)

	b[0][0] = 100
	fmt.Println(a)
	fmt.Println(b)
}
