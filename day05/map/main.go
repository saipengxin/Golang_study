package main

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
)

var m = sync.Map{}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)
			value, _ := m.Load(key)
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
			atomic.AddInt32()
		}(i)
	}
	wg.Wait()
}
