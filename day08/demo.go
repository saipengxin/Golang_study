package main

import "fmt"

const N = 3

func main() {
	m := make(map[int]*int)
	for i := 0; i < N; i++ {
		m[i] = &i
	}

	fmt.Println(m)
}
