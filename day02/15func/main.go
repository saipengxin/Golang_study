package main

import "fmt"

/**
x 参数 int类型
y 参数 int类型
res 返回值 int类型
*/
func sum(x int, y int) (res int) {
	return x + y
}

// 多个返回值
func cheng(x int) (res1 int, res2 int) {
	return x - 1, x - 2
}

func main() {
	// 直接使用函数名调用并使用变量接收返回值
	r := sum(1, 2)
	// 返回多个返回值时，使用多个变量接收
	i, j := cheng(5)
	// 如果存在返回值不需要使用，使用匿名函数来接收
	_, v := cheng(4)
	// 调用有返回值函数的时候，我们也可以不接收返回值
	cheng(6)
	fmt.Println(r)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(v)
}

func intSum(int, int) () {

}
