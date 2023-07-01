package main

import "fmt"

func main() {

	// LABEL为一个标签
	// 以下语句如果break后面不加上标签的话，只会跳出内层循环
	// 但最外层循环是一个死循环，所以此程序会一直执行下去
	// 如果加上标签过后，break就会跳出到跟LABEL同级的循环，这样就可以成功的break
LABEL:
	for {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			if i > 3 {
				break LABEL
			}
		}
	}

	//continue执行就会跳出当前循环，并且继续执行跟标签同级的循环
LABEL1:
	for j := 0; j < 3; j++ {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			if i > 3 {
				continue LABEL1
			}
		}
	}
	//当执行goto过后，就会一直从头开始执行标签，所以程序不会退出，会一直继续下去
LABEL2:
	for j := 0; j < 3; j++ {
		for {
			fmt.Println(j)
			goto LABEL2
		}
	}

}
