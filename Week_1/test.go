package main

func Sum(a, b int) (int, bool, bool) {
	return a + b, a == b, a != b
}
func main() {
	println("Hello Word!")
	print(Sum(1, 2))
}
