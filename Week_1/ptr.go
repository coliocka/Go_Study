package main

import "fmt"

func main() {
	a := 1
	// 在Go语言中不能使用++a这样的式子
	a++
	// 在Go语言中指针是不支持运算的
	var p *int = &a
	fmt.Println(p)
	fmt.Println(*p)

	// 当范围局部变量和局部变量重名时，优先使用范围局部变量，局部变量会被隐藏起来
	// if语句执行完成过后，a = 1
	if a := 10; a > 0 {
		fmt.Println(a)
	}
	fmt.Println(a)
}
