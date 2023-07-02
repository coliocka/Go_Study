package main

import (
	"fmt"
	"sort"
)

func main() {
	// map创建
	// 也可以使用:=进行定义
	var m = map[int]string{}
	fmt.Println(m)
	// 也可以通过make创建
	m = make(map[int]string)
	fmt.Println(m)
	m[1] = "OK"
	fmt.Println(m[1]) // "OK"
	fmt.Println(m[2]) //nui
	delete(m, 1)      // 删除键值对
	fmt.Println(m[1]) //nui

	// 多个map相互嵌套，每一级的map都需要初始化
	// 两个map的嵌套
	mp := map[int]map[int]string{}
	// 上一个语句只对第一个map进行了初始化
	// 必须执行以下这一句，初始化第二层map才能够使用
	mp[1] = make(map[int]string)
	mp[1][1] = "OK"
	// 多返回值，第二个返回值为bool值
	a, ok := mp[2][1]
	fmt.Println(a, ok) // nui false
	// 多返回值的好处就是，如果初始化，可以先初始化，避免报错
	if !ok {
		mp[2] = make(map[int]string)
	}
	mp[2][1] = "Good"
	fmt.Println(mp[2][1])

	// 遍历map
	sm := make([]map[int]string, 5)
	// v是一个拷贝，v的任何操作都不会影响到sm本身
	for i := range sm {
		sm[i] = make(map[int]string, 1)
		sm[i][1] = "OK"
		fmt.Println(sm[i])
	}
	fmt.Println(sm)

	smp := map[int]string{1: "a", 5: "b", 3: "c", 4: "d"}
	s := make([]int, len(smp))

	i := 0
	for k := range smp {
		s[i] = k
		i++
	}
	sort.Ints(s) // 对于int类型进行排序
	fmt.Println(s)

	// 使用for range将类型map[int]string的键和值进行交换，变为map[string]int
	m1 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d"}
	fmt.Println(m1)
	m2 := map[string]int{}
	for k, v := range m1 {
		m2[v] = k
	}
	fmt.Println(m2)
}
