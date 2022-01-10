package main

import (
	"fmt"
	"sync"
)

var b chan int // 定义一个int类型的通道
var wg sync.WaitGroup

// 通道是引用类型，传递过来的就是引用
// 要求ch1只能写入，不能读取
func f1(ch1 chan<- int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

// 通道是引用类型，传递过来的就是引用,能直接修改外部变量
// ch1只能读取，不能写入
// ch2只能写入，不能读取
func f2(ch1 <-chan int, ch2 chan<- int) {
	defer wg.Done()
	for i := range ch1 {
		ch2 <- i * i
	}
	close(ch2)
}

// channel 练习
func main() {
	wg.Add(2)
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 这里定义的是双向通道，但是传递到函数中作为参数转换成了单向通道
	go f1(ch1)
	go f2(ch1, ch2)
	go f2(ch1, ch2)

	for i := range ch2 {
		fmt.Println(i)
	}
	wg.Wait()
}
