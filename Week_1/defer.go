package main

import (
	"fmt"
)

/*
defer 的执行方式类似其他语言中的析构函数，在函数体执行结束后按照调用顺序的相反顺序逐一执行
即使函数发生严重错误也会执行
支持匿名函数的调用
常用于资源清理、文件关闭、解锁以及记录时间等操作
通过与匿名函数配合可在return之后修改函数计算结果
如果函数体内某个变量作为defer时匿名函数的参数，则在定义defer时已经获得了拷贝，否则则是引用某个变量的地址

Go没有异常机制，但有panic/recover模式来处理错误
Panic可以在任何地方引发，但recover只有在defer调用的函数中有效
*/
func main() {
	fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")
	// 逆序向上调用，输出结果为a,c,b

	// 用循环更加直观
	// 这里打印出来是2 1 0
	// 但因为上面还有defer，所以是从最后面的defer开始打印
	// 所以结果就是a 2 1 0 c b
	//for i := 0; i < 3; i++ {
	//	defer fmt.Println(i)
	//}

	for i := 0; i < 3; i++ {
		i := i //如果不在这里进行值的拷贝
		// 输出结果将会全是3
		// 因为defer是在当前函数，也就是main return过后才会逆序调用，
		// defer 使用匿名函数是对i进行了地址的拷贝，所以当循环结束时，i=3，main函数return过后
		// 反过来调用defer，输出的都会是3
		defer func() {
			fmt.Println(i)
		}()
	}

	funcA()
	funcB()
	funcC()

	var fs = [4]func(){}
	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i = ", i)
		defer func() {
			fmt.Println("defer closure i = ", i)
		}()
		fs[i] = func() {
			fmt.Println("closure i = ", i)
		}
	}
	for _, f := range fs {
		f()
	}
}

func funcA() {
	fmt.Println("func_A")
}

func funcB() {
	// recover()会返回一个信息，如果返回不是nil，说明引发panic信息
	defer func() {
		if err := recover(); err != nil {
			println("Recover in B")
		}
	}()
	// panic输出了就终止程序的运行
	// 进行恢复需要进行recover
	panic("panic in func_B")
}

func funcC() {
	fmt.Println("func_c")
}
