package main

import "fmt"

// 结构体
type Options struct {
	strOption1 string
	strOption2 string
	strOption3 string
	intOption1 int
	intOption2 int
	intOption3 int
}

// 声明一个函数类型的变量，用于传参
type Option func(opts *Options)

// 初始化结构体
func InitOptions1(opts ...Option) {
	options := &Options{}
	// 遍历opts，得到每一个函数
	for _, opt := range opts {
		// 调用函数，在函数里，给传进去的对象赋值
		// 也就是说其实赋值的操作不在初始化的代码里，而是在传递过来的函数里。初始化代码的作用是调用传递过来函数。
		opt(options)
	}
	fmt.Printf("init options %#v\n", options)
}

// 定义具体给某个字段赋值的方法
func WithStrOption1(str string) Option {
	return func(opts *Options) {
		opts.strOption1 = str
	}
}

func WithStrOption2(str string) Option {
	return func(opts *Options) {
		opts.strOption2 = str
	}
}
func WithStrOption3(str string) Option {
	return func(opts *Options) {
		opts.strOption3 = str
	}
}
func WithIntOption1(i int) Option {
	return func(opts *Options) {
		opts.intOption1 = i
	}
}
func WithIntOption2(i int) Option {
	return func(opts *Options) {
		opts.intOption2 = i
	}
}
func WithIntOption3(i int) Option {
	return func(opts *Options) {
		opts.intOption3 = i
	}
}

func main() {
	// 我们调用函数赋值
	InitOptions1(WithStrOption1("str1"), WithStrOption3("str3"))
}
