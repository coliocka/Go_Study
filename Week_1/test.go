package main

import "fmt"

func f(p int, t int) int {
	return (p << 1) + (t << 1)
}

func main() {
	fmt.Printf("%d\n", f(10, 10))
}
