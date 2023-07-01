package main

import (
	"fmt"
)

// 第一行必须有常量表达式
// iota计数器，初始值为0，每定义一个常量就会加1
// 并且iota关键字每遇到一个const关键字，就会重置为0
const (
	A = "A"
	B // 这里常量表达式未定义默认使用上一个常量表达式的值
	C = iota
	D // 这里默认使用上一个常量表达式的值，因为这定义了第四个常量，所以d = iota = 3
)

// 这里iota又变为了0
const (
	E = 10
	F = iota
	G
	H
)

func main() {
	fmt.Println(^100)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
	fmt.Println(G)
}
