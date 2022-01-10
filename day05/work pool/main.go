package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

type job struct {
	x int64
}

type result struct {
	job *job
	sum int64
}

var jobChan = make(chan *job, 100)
var resultChan = make(chan *result, 100)

func sheng(jobChan chan<- *job) {
	defer wg.Done()
	for {
		// 生成随机数
		x := rand.Int63()
		j := &job{
			x,
		}
		jobChan <- j
		time.Sleep(time.Second)
	}
}

func jisuan(jobChan <-chan *job, resultChan chan<- *result) {
	defer wg.Done()
	for {
		job := <-jobChan
		var sum int64 = 0
		n := job.x
		for n > 0 {
			sum += (n % 10)
			n = n / 10
		}
		r := &result{job: job, sum: sum}

		resultChan <- r
	}
}

func main() {
	wg.Add(1)
	go sheng(jobChan)

	for i := 0; i < 24; i++ {
		wg.Add(1)
		go jisuan(jobChan, resultChan)
	}

	for res := range resultChan {
		fmt.Printf("job:%d,sum:%d\n", res.job.x, res.sum)
	}
	wg.Wait()
}
