package protocol

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

// Encode 将消息编码，携带数据长度发送,采用固定4字节存储
func Encode(message string) ([]byte, error) {
	// 读取消息的长度，转换成int32类型（占4个字节）
	var length = int32(len(message))
	// 创建一个缓冲区
	// Buffer是一个实现了读写方法的可变大小的字节缓冲。本类型的零值是一个空的可用于读写的缓冲。
	var pkg = new(bytes.Buffer)

	// 采用小端的方式写入消息头(无论大端小端，编码和解码统一即可)
	// 将length的binary编码格式写入pkg，length必须是定长值、定长值的切片、定长值的指针。binary.LittleEndian指定写入数据的字节序，写入结构体时，名字中有'_'的字段会置为0。
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	// 写入消息实体
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	// 返回未读取部分字节数据的切片，len(b.Bytes()) == b.Len()。如果中间没有调用其他方法，修改返回的切片的内容会直接改变Buffer的内容。
	return pkg.Bytes(), nil
}

// Decode 解码消息
func Decode(reader *bufio.Reader) (string, error) {
	// 读取消息的长度
	// Peek返回输入流的下n个字节，而不会移动读取位置。返回的[]byte只在下一次调用读取操作前合法。
	lengthByte, _ := reader.Peek(4) // 读取前4个字节的数据

	// NewBuffer使用指定内容作为初始内容创建并初始化一个Buffer
	lengthBuff := bytes.NewBuffer(lengthByte)

	var length int32
	err := binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}
	// Buffered返回缓冲中现有的可读取的字节数。
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}
	// 读取真正的消息数据
	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}
