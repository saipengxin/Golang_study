package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	UserName string `ini:"username"`
	PassWord string `ini:"password"`
}

type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database string `ini:"database"`
}

type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func main() {
	var cfg Config
	// 第二个参数这里一定是传递指针，因为要从函数内修改函数外的这个变量里的值，如果不是指针，修改的是副本
	// 而且如果不是指针反射赋值的时候会panic
	err := loadIni("./demo.ini", &cfg)
	if err != nil {
		fmt.Println("解析失败：", err)
		return
	}
	fmt.Printf("%#v", cfg)
}

func loadIni(fileName string, data interface{}) (err error) {
	// 1、参数校验
	// 1.1 确保是data指针类型，否则反射赋值会panic，而且不是指针函数内修改的是变量副本，不是原变量
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	if t.Kind() != reflect.Ptr {
		err = errors.New("变量类型错误，请确保是指针类型")
		return
	}
	// 1.2 确保data的值是结构体类型，因为这里要将配置文件中的值赋值给结构体中的字段,也就是说data不光是一个指针，还要是一个结构体指针
	// 指针在反射中，使用Elem函数来获取指向的值，我们要确定他的值是结构体类型
	if v.Elem().Kind() != reflect.Struct {
		err = errors.New("变量类型错误,请确保是结构体类型")
		return
	}

	// 2、读文件得到字节类型数据，这里选择的方式是将文件内容整体读出来保存在变量中，然后遍历这个变量
	// 因为如果我们打开文件一行行的读，可能会因为程序运行缓慢导致文件句柄被打开好久才能关闭。
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		err = errors.New("文件读取失败")
		return
	}
	// 返回的是字节切片类型数据，我们要将他转为字符串处理
	// 将数据转为字符串后，根据换行来将内容分割成切片，因为配置项都是一行一个的
	lineSlice := strings.Split(string(b), "\n")
	var structName string
	for idx, line := range lineSlice {
		// 首先去除这一行中首尾的空格，这个很重要，否则我们判断内容的时候无法确定第一个字符是否是 [ 开头还是内容开头，很有可能取到空格
		// 去除首尾空格 strings.TrimSpace()
		line = strings.TrimSpace(line)

		// 我们可能会将不同配置项的内容之间添加空行来分割，空行不需要进行处理，直接跳过
		if len(line) == 0 {
			continue
		}

		// 注释也要进行跳过,ini文件 注释是 # 或者 ; 开头的行
		// strings.HasPrefix():是否以某字符串开头
		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, ";") {
			continue
		}

		// 如果是[开头的就表示是节(section)
		if strings.HasPrefix(line, "[") {
			//判断格式是否正确，因为我们要根据这个来判断是那个配置项的值，添加到那个结构体中 例如 [mysql]
			// 保证开头是 [ , 结尾是 ]
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("行 %d 格式错误", idx+1)
				return
			}

			// 保证 [ ] 中间有内容
			// 字符串本质是只读型的字符数组，所以我们可以使用下面的方式来进行字符串的切割，使用方式和数组、切片相同
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("行 %d ,语法错误", idx+1)
				return
			}

			// 取到正确的节点名称，我们要去结构体指针中找到对应的子结构体元素，准备开始赋值
			// 由于我们无法确定字段名称和tag中的名称对应关系，所以要一个一个变量的去找
			for i := 0; i < v.Elem().NumField(); i++ {
				// 遍历每一个字段，获取到他的ini对应的tag值，然后和节点名称比较
				field := t.Elem().Field(i) // 字段的tag信息，存在类型对象中 TypeOf
				if sectionName == field.Tag.Get("ini") {
					// 相等，说明找到了对应的嵌套结构体，将这个结构体名称记录下来，下一遍循环键值对的时候，要向这个结构体中的字段赋值
					structName = field.Name
				}
			}
		} else {
			// 如果不是 [ 开头就是=分割的键值对

			// 如果简直对中没有等号或者   =3306  这种格式都是不对的。所以要判断是否存在等号并且不是等号开头，等号结尾看情况而定，可以当作空值
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("行 %d ,语法错误")
				return
			}

			// 获取到等号所在位置，并根据等号位置取到键名和键值
			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])
			value := strings.TrimSpace(line[index+1:])

			// 根据结构体名称，去指针中取到对应结构体的相关信息
			sValue := v.Elem().FieldByName(structName) // 拿到嵌套的结构体值信息
			sType := sValue.Type()                     // 拿到值对应的类型信息
			// 如果嵌套的不是结构体要报错
			if sType.Kind() != reflect.Struct {
				err = fmt.Errorf("%s应该是一个结构体", structName)
				return
			}

			var fieldName string
			// 遍历嵌套结构体，判断字段的tag是否和key相等
			for i := 0; i < sValue.NumField(); i++ {
				field := sType.Field(i) // tag信息是存在类型对象中的
				if field.Tag.Get("ini") == key {
					// 记录字段名称，赋值的时候使用
					fieldName = field.Name
					break
				}
			}
			// 在结构体中没有找到字段信息，跳过
			if len(fieldName) == 0 {
				continue
			}
			// 取到嵌套结构提中的字段的相关信息
			fieldObj := sValue.FieldByName(fieldName)
			// 对取到的字段信息赋值，赋值要根据类型来赋值
			switch fieldObj.Type().Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				// 将字符串值转为数值型
				var valueInt int64
				valueInt, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					err = fmt.Errorf("行:%d 值类型错误", idx+1)
					return
				}
				fieldObj.SetInt(valueInt)
			case reflect.String:
				fieldObj.SetString(value)
			case reflect.Bool:
				var valueBool bool
				valueBool, err = strconv.ParseBool(value)
				if err != nil {
					err = fmt.Errorf("行:%d 值类型错误", idx+1)
					return
				}
				fieldObj.SetBool(valueBool)
			case reflect.Float32, reflect.Float64:
				var valueFloat float64
				valueFloat, err = strconv.ParseFloat(value, 64)
				if err != nil {
					err = fmt.Errorf("行:%d 值类型错误", idx+1)
					return
				}
				fieldObj.SetFloat(valueFloat)
			}
		}
	}

	return nil
}
