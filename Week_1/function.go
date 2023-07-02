package main

import "fmt"

func main() {
	fmt.Println(fc(1, 2, 3, "o"))
	fun(1, 2, 3, 4)
	a, b := 1, 2
	function1(a, b)
	fmt.Println(a, b)
	s := []int{1, 2, 3}
	Function2(s)
	fmt.Println(s)

	c := 3
	function3(&c, b)
	fmt.Println(c, b)

	// 匿名函数
	d := func() {
		fmt.Println("hello")
	}
	d()

	f := closure(10)
	fmt.Println(f(10))
}

func fc(a, b, c int, d string) (int, string, int) {
	a, b, c = 1, 2, 3
	d = "hello"
	return a, d, c
}

// 不定长参数使用方法，不定长变参只能放在参数列表的最后面
// 例如：fun(a ...int, b string) 这样是不可以的
// 只能这样 fun(b string, a ...int)
// 传进来过后就是一个slice
func fun(a ...int) {
	fmt.Println(a)
}

// slice在使用的时候传递的是引用，但是使用不定长变参的时候，传递进来会变成一个slice
// 但在函数内对不定长变参传进来的slice进行改变，不会对函数外传进来的数字进行改变
func function1(s ...int) {
	s[0], s[1] = 3, 4
	fmt.Println(s)
}

// Function2 可以看到，当传递过来的数不是不定长变参，而是传递slice，就可以对原有的slice进行修改
// 是对slice地址的拷贝
func Function2(s []int) {
	s[0] = 10
	fmt.Println(s)
}

// 如果想要对传进来的数进行修改的话，必须使用指针传参，进行地址拷贝，而不是值拷贝
// 可以看到如果不进行指针传参，也就是地址拷贝的话，b的值在函数中进行修改是无法改变函数外参数的值的
func function3(a *int, b int) {
	*a, b = 100, 10
	fmt.Println(*a, b)
}

// 闭包，返回一个匿名函数
// 调用的x是直接指向x的地址，而不是值拷贝
func closure(x int) func(int) int {
	fmt.Println("FUNC closure")
	fmt.Printf("%p\n", &x)
	return func(y int) int {
		fmt.Printf("%p, %p\n", &x, &y)
		return x + y
	}
}
