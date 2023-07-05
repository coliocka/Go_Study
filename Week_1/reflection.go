package main

import (
	"fmt"
	"reflect"
)

/*
反射reflection
反射可大大提高程序的灵活性，使得interface{}有更大的发挥余地
反射使用TypeOf和ValueOf函数从接口中获取目标对象信息
反射会将匿名字段作为独立字段（匿名字段本质）
想要利用反射修改对象状态，前提是interface.data是settable，即pointer-interface
通过反射可以“动态”调用方法
*/

type User struct {
	Id   int
	Name string
	Age  int
}

// Manager 反射会将匿名字段当成独立字段来处理
type Manager struct {
	User
	title string
}

func (u User) Hello(name string) {
	fmt.Println("Hello, ", name, "My name is: ", u.Name)
}

func main() {
	u := User{1, "OK", 12}
	Info(u)

	m := Manager{User: User{1, "OK", 3}, title: "123123"}
	t := reflect.TypeOf(m)

	fmt.Printf("%#v\n", t.Field(0))
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 1}))

	// 通过反射修改内容
	x := 123
	v := reflect.ValueOf(&x)
	v.Elem().SetInt(999)
	fmt.Println(x)

	u.Hello("joe")

	Set(&u)
	fmt.Println(u)
}

func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())

	//fmt.Println(t.Kind())
	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("Error")
		return
	}

	// v为当前参数o的值
	v := reflect.ValueOf(o)

	fmt.Println("Fields")
	// t为取出当前o的Type
	// Type的NumField就是当前参数值所含几个字段
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}

	// 当前o所有的Method，也就是方法字段
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type)
	}
}

// Set 对于一个对象，如何通过它的反射来进行值的修改
func Set(o interface{}) {
	// 对于所有反射的用法，都需要取出其对应的ValueOf
	v := reflect.ValueOf(o)

	// v.kind()获取的是v的类型
	// 判断是否指针，并且是否可以被修改，如果不能被修改，就输出XXX并返回
	// 如果可以被修改，就通过v.Elem()方法获取指针指向的值，并将其赋值给v变量
	if k := v.Kind(); k == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("XXX")
		return
	} else {
		v = v.Elem()
	}

	// 获取结构体中名为Name的字段，如果不存在，输出Bad并返回
	f := v.FieldByName("Name")
	if !f.IsValid() {
		fmt.Println("Bad")
		return
	}
	// 如果当前类型是string
	// 就将其值设置为ByeBye
	if f.Kind() == reflect.String {
		f.SetString("ByeBye")
	}
}
