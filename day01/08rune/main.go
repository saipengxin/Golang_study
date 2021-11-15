package main

import "fmt"

func main()  {
	var i1 byte = 'a'
	var i2 uint8 = 'b'
	fmt.Println(i1)
	fmt.Println(i2)
	fmt.Printf("%c\n",i1)
	fmt.Printf("%c\n",i2)

	var i3 rune = '赛'
	var i4 int32 = '赛'
	var i5 int64 = '赛'
	fmt.Println(i3)
	fmt.Println(i4)
	fmt.Println(i5)
	fmt.Printf("%c\n",i3)
	fmt.Printf("%c\n",i4)
	fmt.Printf("%c\n",i5)
	fmt.Println(string(i5))
}
