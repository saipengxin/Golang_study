package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var x int32 = 0
var wg sync.WaitGroup

func f1() {
	defer wg.Done()
	for i := 0; i < 5000; i++ {
		atomic.AddInt32(&x, 1)
	}
}

func main() {
	wg.Add(2)
	go f1()
	go f1()
	wg.Wait()
	fmt.Println(x)
}
