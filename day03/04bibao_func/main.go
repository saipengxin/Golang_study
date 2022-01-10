package main

import "fmt"

<<<<<<< HEAD
func f1() func(int) {
	var x = 10
	return func(y int) {
		x += y
		fmt.Printf("%p\n", x)
	}
}

func f2(y int) {
	var x = 10
	x = x + y
	fmt.Printf("%T\n", x)
}

func main() {
	var f = f1()
	f(10)
	f(20)
	fmt.Println("======================")
	var ff = f1()
	ff(10)
	ff(20)
	fmt.Println("======================")
	var fff = f2
	fff(10)
	fff(20)
=======
/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
}

func dispatchCoin() int {
	for _, v := range users {

	}
>>>>>>> 1fddb6ff01839fec3e0d5bdf8193ea5d03003eac
}
