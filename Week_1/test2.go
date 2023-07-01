package main

import (
	"fmt"
)

/*
6 	0110
11 	1011
************
&   0010 = 2 	同1为1，否则为0
^	1101 = 13	同1为0，否则为1
|	1111 = 15	同0为0，否则为1
&^	0100 = 4	当前位第二个数为1，第一个数当前位就为0，其余不变
***********
*/
// _ 是一个特殊的标识符，成为空标识符，表示一个不需要使用的变量或值
// iota自增，就可以实现对KB, MB, GB, TB 进行赋值
const (
	_          = iota
	KB float64 = 1 << (iota * 10)
	MB
	GB
	TB
)

func main() {
	fmt.Println(6 & 11)
	fmt.Println(6 ^ 11)
	fmt.Println(6 | 11)
	fmt.Println(6 &^ 11)
	a := 0
	fmt.Printf("KB=%fB\n", KB)
	fmt.Printf("MB=%fB\n", MB)
	fmt.Printf("GB=%fB\n", GB)
	fmt.Printf("TB=%fB\n", TB)
	// &&运算符表示左边的表达式返回为true才会执行右边的表达式
	if a > 0 && (10/a > 1) {
		fmt.Println("OK!")
	}

	b := 10
	// ||运算符表示两边任意一个表达式返回为true就继续执行括号内的语句
	if b > 11 || (10/b) >= 1 {
		fmt.Println("B OK!")
	}
}
