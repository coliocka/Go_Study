package main

import (
	"fmt"
)

type T struct {
}

func init() { // 初始化包

}

func main() {
	var a int
	Func1()
	var t T
	t.Method1()

	fmt.Println(a)
}

// （t T）表示该函数是定义在类型T上的方法，该方法的接收者为t
// 通常，方法定义中的接收者类型可以是任何类型，包括基本类型，结构体类型，指针类型等
// 通过在方法上定义接收者，可以将该方法与接收者类型绑定在一起，从而实现了一种面对对象编程的方式
// 例如，在某个类型上定义一个方法，可以使该类型具备某种特定的行为或功能，从而让代码更加清晰、简洁和易于维护
func (t T) Method1() {
	fmt.Println("Method1")
}

func Func1() {
	fmt.Println("Func1")
}
