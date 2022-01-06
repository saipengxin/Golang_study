package sum

import "fmt"

func init() {
	fmt.Println("pach02的calc中的init")
}

func Calc(x, y int) int {
	return x - y
}
