package main

import (
	"fmt"
	"strconv"
)

/*
接口interface
接口是一个或多个方法签名的集合
只要某个类型拥有该接口的所有方法签名，即算实现该接口，无需显示声明实现了哪个接口，这称为Structural Typing
接口只有方法声明，没有实现，没有数据字段
接口可以匿名嵌入其他接口，或嵌入到结构中
将对象赋值给接口时，会发生拷贝，而接口内部存储的是指向这个复制品的指针，既无法修改复制品的状态，也无法获取指针
只有当接口存储的类型和对象都为nil时，接口才等于nil
接口调用不会做receiver的自动转换
接口同样支持匿名字段方法
接口也可实现类似OOP中的多态
空接口可以作为任何类型大数据的容器
*/

// USB 封装了两个接口
type USB interface {
	Name() string
	connect
}

type connect interface {
	Connect()
}

type Phone struct {
	name string
}

// Name 当前接口属于Phone
func (pc Phone) Name() string {
	return pc.name
}

// Connect 当前接口属于Phone
func (pc Phone) Connect() {
	fmt.Println("Connect to Phone", pc.name)
}

type Phone1 struct {
	num int
}

func (pc Phone1) Name() string {
	return strconv.Itoa(pc.num)
}

func (pc Phone1) Connect() {
	fmt.Println("Connect to Phone_1", pc.num)
}

func main() {
	var a USB
	a = Phone{"lds"}
	a.Connect()
	fmt.Println(a.Name())

	b := Phone1{12}
	b.Connect()
	fmt.Println(b.num)
	Disconnect(a)
}

// Disconnect 空接口可以传递任意东西
// 当参数是接受空接口的时候，用type switch方法比较多
func Disconnect(usb interface{}) {
	//if pc, ok := usb.(Phone); ok {
	//	fmt.Println("Disconnected:", pc.name)
	//	return
	//}

	switch v := usb.(type) {
	case Phone:
		fmt.Println("Disconnected:", v.name)
	default:
		fmt.Println("Unknown device")
	}
}
