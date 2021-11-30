package main

import "fmt"

func main()  {
	var s1 = []int{1,2,3,4,5}
	for i := 0; i < len(s1); i++ {
		fmt.Println(s1[i])
	}

	fmt.Println("=====================")

	for _,v := range s1 {
		fmt.Println(v)
	}
}
