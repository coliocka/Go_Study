package main

import "fmt"
import "math/rand"
import "time"

// 随机数组
func initData(arr []int) {
	for i := range arr {
		arr[i] = rand.Intn(100)
	}
}

func QuickSort(arr []int, index_start int, index_end int) {
	//递归函数出口
	if index_start >= index_end {
		return
	}
	//基准数
	num, i, j := arr[index_start], index_start, index_end
	//找到num合适的位置
	for i < j {
		//将比基准数小的放到它的左边
		for i < j {
			if arr[j] < num {
				arr[i] = arr[j]
				i++ //i往前移动
				break
			}
			j--
		}
		//将比基准数大的放到它的右边
		for i < j {
			if arr[i] >= num {
				arr[j] = arr[i]
				j-- //将往后移动
				break
			}
			i++
		}
	}
	//将基准数放到 它应该有的位置
	arr[i] = num
	//将左边 进行同样的排序
	QuickSort(arr, index_start, i)
	//将右边 进行同样的排序
	QuickSort(arr, i+1, index_end)
}

func main() {
	//设置随机数种子
	rand.Seed(time.Now().UnixNano())

	n := 10
	arr := make([]int, n, n)
	//初始化切片数组
	initData(arr)
	fmt.Println(arr)
	//快速排序
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
