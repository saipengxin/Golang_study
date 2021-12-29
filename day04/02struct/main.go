package main

import "fmt"

func main() {
	var a = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	// 遍历a
	for i, v := range a {
		fmt.Println(v)
		fmt.Printf("%p\n", &v)
		fmt.Printf("%p\n", &a[i])
		fmt.Println("============")
	}

}
