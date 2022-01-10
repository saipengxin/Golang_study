package main

import (
	"fmt"
	"time"
)

func main() {
	var time_str = "2022/01/01 10:20:30"

	timeObj, err := time.Parse("2006/01/02 15:04:05", time_str)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj.Unix())
}
