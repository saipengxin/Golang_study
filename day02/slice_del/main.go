package main
import "fmt"
func main()  {
	var a = [8]int{1,2,3,4,5,6,7,8} // 定义底层数组
	b := a[1:4] // 使用自定义的底层数组转为切片 [2 3 4]

	c := b[5:5] // 切片再切片
	fmt.Println(c)
}