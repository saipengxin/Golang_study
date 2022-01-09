package ref

import (
	"fmt"
	"reflect"
)

// Check_action ...
func Check_action(a interface{}) bool {
	// 首先判断这是不是一个结构体对象
	if k := reflect.TypeOf(a).Kind(); k != reflect.Struct {
		fmt.Println("类型错误")
		return false
	}

	// 判断结构体对象中是否存在action方法
	fmt.Println()
	if !reflect.ValueOf(a).MethodByName("Action").IsValid() {
		fmt.Println("请定义Action方法")
		return false
	}

	// 判断是否存在前置方法 _before_action
	if reflect.ValueOf(a).MethodByName("Before_action").IsValid() {
		// 如果存在执行自动执行
		reflect.ValueOf(a).MethodByName("Before_action").Call([]reflect.Value{})
	}

	// 执行action
	reflect.ValueOf(a).MethodByName("Action").Call([]reflect.Value{})

	// 判断是否存在后置方法 _after_action
	if reflect.ValueOf(a).MethodByName("After_action").IsValid() {
		reflect.ValueOf(a).MethodByName("After_action").Call([]reflect.Value{})
	}

	return true
}
