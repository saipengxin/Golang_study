package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := `aaaa

赛`
	fmt.Println(strconv.Quote(a))
	fmt.Println(strconv.QuoteToASCII(a))

	b := '赛'
	fmt.Println(strconv.QuoteRune(b))
	fmt.Println(strconv.QuoteRuneToASCII(b))

	c := "'赛'"
	fmt.Println(strconv.Unquote(c))
	d := "\"fdsafsda\""
	fmt.Println(strconv.Unquote(d))
	e := "'abc'"
	fmt.Println(strconv.Unquote(e))
}
