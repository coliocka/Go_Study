package main

import (
	. "fmt"
)

// 匿名结构嵌套
type person struct {
	Name    string
	Age     int
	Contact struct {
		Phone, City string
	}
}

// 组合
type human struct {
	Sex int
}

type teacher struct {
	human
	Name string
	Age  int
}

type student struct {
	human
	Name string
	Age  int
}

func main() {
	a := &person{
		Name: "lds",
		Age:  19,
	}
	a.Contact.Phone, a.Contact.City = "1231312", "chengdu"

	Println(a)
	fA(a)
	Println(a)

	// 匿名结构
	b := &struct {
		Name string
		Age  int
	}{
		Name: "joe",
		Age:  12,
	}
	Println(b)

	c := &teacher{Name: "lds", Age: 20, human: human{Sex: 0}}
	d := &student{Name: "lds", Age: 11, human: human{Sex: 1}}
	c.Name = "joe"
	c.Age = 14
	c.Sex = 2
	Println(c, d)
}

// 函数中也是得到一个值的拷贝
func fA(per *person) {
	per.Age = 13
	Println("A", per)
}
