package main

import (
	"fmt"
	"sync"
)

var x = 0
var wg sync.WaitGroup
var lock sync.Mutex

func f1() {
	defer wg.Done()
	for i := 0; i < 5000; i++ {
		lock.Lock() // 加锁，保证同一时刻只有一个goroutine来操作变量
		x = x + 1
		lock.Unlock() // 解锁，不解锁后面会阻塞
	}

}

func main() {
	wg.Add(2)
	go f1()
	go f1()
	wg.Wait()
	fmt.Println(x)
}
