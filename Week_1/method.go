package main

import (
	"fmt"
)

type TZ int

type AA struct {
	Name string
}

type BB struct {
	Name string
}

func main() {
	a := AA{}
	a.Print()
	fmt.Println(a.Name) //AA

	b := BB{}
	b.Print()
	fmt.Println(b.Name) //nil

	var c TZ
	c.Print()

	(*TZ).Print(&c)

	var d TZ
	d.Increase(100)
	fmt.Println(d)
}

func (b *TZ) Increase(num int) {
	*b += TZ(num)
}

// Print 属于A的方法
func (a *AA) Print() {
	a.Name = "AA"
	fmt.Println("A")
}

// Print 所以遵循参数规则
// 以指针形式传递，是地址的拷贝
// 以变量的形式传递，是值拷贝
func (b BB) Print() {
	b.Name = "BB"
	fmt.Println("B")
}

// Print 结构绑定在TZ上
func (b *TZ) Print() {
	fmt.Println("TZ")
}
