package main

import (
	. "fmt"
)

func main() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	Println(a)

	// 左开右闭的区间，也就是切a[5],a[6],a[7],a[8],a[9]
	s1 := a[5:10]
	Println(s1)

	// 切出后5个元素
	s2 := a[5:]
	Println(s2)

	// 切出前5个元素
	s3 := a[:5]
	Println(s3)

	// make构造切片，第一参数为切片所指向的类型，第二个参数为当前容量，第三个参数为申请的最大容量
	// 当然也可以不设置最大容量，编译器就会自动认为当前最大容量就是当前容量
	s4 := make([]int, 10, 10)
	s4 = append(s4, a[0])
	Println(len(s4), cap(s4)) // 11, 20
	for i := 0; i < 10; i++ {
		s4 = append(s4, 3)
	}
	Println(s4)
	// 可以看到当最大容量已经充满时，如果我们要再次加入元素，是可以的，但编译器会自动帮我们再申请一次跟最大容量等价的元素个数
	Println(len(s4), cap(s4)) // 11,  40

	u := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}
	saU := u[2:5]
	// 可以看出，len(sa_u)输出为3，正好是切片sa_u的长度，但是cap是6
	// 就是切片对象u从下标为2的元素开始，一直到最后的长度，底层的最大容量
	// 索引不可以超过被slice的切片的容量cap()值
	// 索引越界不会导致底层数组的重新分配而是引发错误
	Println(len(saU), cap(saU)) // 3, 6
	Println(string(saU))
	Println(string(u))
	// 从切片中构建切片
	sbU := saU[:3]
	Println(string(sbU))

	// Append
	// 可以在slice尾部追加元素
	// 可以将一个slice追加在另一个slice尾部
	// 如果最终长度未超过追加到slice的容量则返回原始slice
	// 如果超过追加到的slice的容量则将重新分配数组并拷贝原始数据
	s1 = make([]int, 3, 6)
	Printf("%p\n", s1) // 打印切片的地址
	Println(s1)
	s1 = append(s1, 1, 2, 3)
	Printf("%p\n", s1) // 没有重新分配容量，地址还是相同
	Println(s1)
	s1 = append(s1, 1, 2, 3)
	// 当当前需要容量大于原始最大容量时，就会重新分配容量，再将原始数据拷贝到当前内存里
	Printf("%p\n", s1) // 重新分配容量，地址也发生了改变
	Println(s1)

	a1 := []int{1, 2, 3, 4, 5}
	sA := a1[2:5]
	sB := a1[1:3]
	Println(sA, sB) // 3, 4, 5    2, 3
	// 当append超过底层最大容量时，slice就会指向一个新的底层数组，不再指向a1
	// 所以当s_a去改变数组a1时，不会影响s_b
	sB = append(sB, 1, 2, 3, 1, 1, 1, 1, 1, 1, 1, 1)
	sA[0] = 9
	// 说明slice是指向底层的数组，它改变也会改变数组
	Println(sA, sB, a1) //[9 4 5] [2 3 1 2 3 1 1 1 1 1 1 1 1] [1 2 9 4 5]

	// copy
	slice1 := []int{1, 2, 3, 4, 5, 6}
	slice2 := []int{7, 8, 9}
	//copy(slice_1, slice_2)
	// slice_2中的值拷贝到slice_1中，所以只改变了slice_1中前三个值
	//Println(slice_1, slice_2) //[7 8 9 4 5 6] [7 8 9]
	copy(slice2[2:3], slice1[1:3])
	// 就是把slice_1中的索引为1,2,3的元素拷贝到slice_2中索引为2的元素上，因为slice_2中只有一个需要拷贝的
	// 所以就只拷贝slice_1中索引为1的元素，也就是2，到slice_2[2]上
	Println(slice1, slice2) //[1 2 3 4 5 6] [7 8 2]

	slice3 := slice1[:cap(slice1)] // 将slice_1切片赋值给slice_3切片
	slice4 := slice1[:]
	Println(slice3)
	Println(slice4)
}
