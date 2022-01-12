package main

// 大端 小端

import (
	"encoding/binary"
	"fmt"
)

func BigEndian() { // 大端序
	// 二进制形式：0001 0010 0011 0100 0101 0110 0111 1000
	var testInt int32 = 0x12345678 // 十六进制表示
	fmt.Printf("%d use big endian: \n", testInt)

	var testBytes []byte = make([]byte, 4)
	// BigEndian:大端，binary包中定义好的
	// PutUint32将uint32类型的数字转为字节序列,uint32占4字节，一字节是8个二进制位
	binary.BigEndian.PutUint32(testBytes, uint32(testInt)) //大端序模式
	fmt.Println("int32 to bytes:", testBytes)

	//Uint32将字节序列转为uint32类型的数字
	convInt := binary.BigEndian.Uint32(testBytes) //大端序模式的字节转为int32
	fmt.Printf("bytes to int32: %d\n\n", convInt)
}

func LittleEndian() { // 小端序
	//二进制形式： 0001 0010 0011 0100 0101 0110 0111 1000
	var testInt int32 = 0x12345678 // 16进制
	fmt.Printf("%d use little endian: \n", testInt)

	var testBytes []byte = make([]byte, 4)
	// LittleEndian:小端，binary包中定义好的
	binary.LittleEndian.PutUint32(testBytes, uint32(testInt)) //小端序模式
	fmt.Println("int32 to bytes:", testBytes)

	convInt := binary.LittleEndian.Uint32(testBytes) //小端序模式的字节转换
	fmt.Printf("bytes to int32: %d\n\n", convInt)
}

func main() {
	BigEndian()
	LittleEndian()
}
