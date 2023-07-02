package main

import "fmt"

func main() {
	var a, b [2]int
	c := [20]int{19: 1}  // 将索引为19的元素设置为1
	d := [...]int{19: 1} //尽可能包含所有数组，所以d数组的大小为20
	var p *[20]int = &c
	fmt.Println(p) // 打印结果多了一个&符号，说明是对c取地址
	fmt.Println(c)
	fmt.Println(d)
	for i := 0; i < 2; i++ {
		fmt.Println(a[i], b[i])
	}

	// 指针数组
	// 打印出来是x,y分别对应的地址
	x, y := 1, 2
	f := [...]*int{&x, &y}

	fmt.Println(f)

	// 用new关键字创建指针数组，返回的是一个指向数组的指针
	q := new([10]int)
	t := [10]int{}
	t[1] = 2
	fmt.Println(t)
	q[1] = 2
	fmt.Println(q)

	//多维数组
	k := [2][3]int{
		{1, 1, 1},
		{2, 3, 2}}
	fmt.Println(k)

	//冒泡排序
	op := [...]int{4, 5, 2, 1, 6}
	fmt.Println(op)

	// 降序
	n := len(op)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if op[i] < op[j] {
				tmp := op[i]
				op[i] = op[j]
				op[j] = tmp
			}
		}
	}
	fmt.Println(op)
}
