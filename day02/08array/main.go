package main

import "fmt"

func main()  {
	var arr = [5]int8{1, 3, 5, 7, 8}
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[i]+arr[j]==8 {
				fmt.Println(i,j)
			}
		}
	}
}
