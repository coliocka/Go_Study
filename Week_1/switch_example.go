package main

import "fmt"

// switch的用法也有三种
func main() {
	a := 1
	switch a {
	case 0:
		fmt.Println("a=0")
	case 1:
		fmt.Println("a=1")
	default:
		fmt.Println("None")
	}

	// 当前a=1，不加fallthrough的话，执行了第一句就跳出了
	// 如果想继续执行的话，必须加上fallthrough
	switch {
	case a >= 0:
		fmt.Println("a>=0")
		fallthrough
	case a >= 1:
		fmt.Println("a>=1")
	case a == 1:
		fmt.Println("a==1")
	default:
		fmt.Println("None")
	}

	// 当然switch语句后面也是可以定义变量的
	switch b := 1; {
	case b >= 0:
		fmt.Println("b>=0")
		fallthrough
	case b >= 1:
		fmt.Println("b>=1")
	case b == 1:
		fmt.Println("b==1")
	default:
		fmt.Println("None")
	}
}
