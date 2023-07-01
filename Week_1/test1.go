package main

func sum(a, b int) (int, bool, bool) {
	return a + b, a == b, a != b
}
func main() {
	println("Hello Word!")
	print(sum(3, 3))
}
