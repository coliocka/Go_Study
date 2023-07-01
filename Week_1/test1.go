package main

// 与c++不同，函数可以返回多个值
func sum(a, b int) (int, bool, bool) {
	return a + b, a == b, a != b
}
func main() {
	println("Hello Word!")
	print(sum(3, 3))
}
