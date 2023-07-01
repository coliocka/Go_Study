package main

import "fmt"

// Go中只有一个循环，就是for，以下演示for的三种常用方法
func main() {
	a := 1
	for {
		a++
		if a > 3 {
			break
		}
		fmt.Println(a)
	}
	fmt.Println(a)
	fmt.Println("over")

	b := 1
	for b <= 3 {
		b++
		fmt.Println(b)
	}
	fmt.Println(b)
	fmt.Println("over")

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
}
