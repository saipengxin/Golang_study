package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

func main() {
	fileName := "./my.log"
	config := tail.Config{
		ReOpen:    false,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}

	tails, err := tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail file err : ", err)
		return
	}
	var (
		msg *tail.Line
		ok  bool
	)
	for {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("msg:", msg.Text)
	}
}
