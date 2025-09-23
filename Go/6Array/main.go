package main

import (
	"fmt"
	"sort"
)

func main() {
	// //从[1,3,5,7,8]中取出两个数字 使其和为8
	// var arr = [...]int{1, 3, 5, 7, 8}
	// for i := 0; i < len(arr); i++ {
	// 	for j := i + 1; j < len(arr); j++ {
	// 		if arr[i]+arr[j] == 8 {
	// 			fmt.Printf("找到一对数字: %d 和 %d\n", arr[i], arr[j])
	// 		}
	// 	}
	// }

	//切片是引用类型
	//切片只声明没赋值的话就是一个nil
	// var arr1 []int
	// fmt.Println(arr1)

	//基于数组定义切片
	// a := [5]int{55, 56, 57, 58, 59} //这是个数组
	// b := a[:]                       //获取数组中的所有值
	// fmt.Printf("%v--%T\n", b, b)
	// c := a[1:4] //左闭右开
	// fmt.Printf("%v--%T\n", c, c)
	//基于切片定义切片 跟上面一样

	//关于长度和容量
	//长度是指切片当前拥有的元素个数
	//容量是指从切片的第一个元素开始，到底层数组末尾的所有元素的数量
	// s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// a := s[2:]
	// fmt.Printf("长度%d 容量%d\n", len(a), cap(a))
	// b := s[1:3]
	// fmt.Printf("长度%d 容量%d\n", len(b), cap(b))

	//make函数创建切片
	//make([]T,len,cap)
	// a := make([]int, 4, 8)
	// fmt.Printf("长度%d 容量%d\n", len(a), cap(a))

	//golang中不能直接用下标的方式给切片扩容
	//扩容需要用append()方法
	// sliceA := []int{1, 23, 4, 5}
	// fmt.Printf("长度%d 容量%d\n", len(sliceA), cap(sliceA))
	// sliceA = append(sliceA, 6, 7, 8, 9)
	// fmt.Printf("长度%d 容量%d\n", len(sliceA), cap(sliceA))

	//切片的扩容策略
	// sliceA := []int{}
	// for i := 1; i < 10; i++ {
	// 	sliceA = append(sliceA, i)
	// 	fmt.Printf("%v 长度:%d 容量:%d\n", sliceA, len(sliceA), cap(sliceA))
	// }

	//选择排序
	// var numSlice = []int{9, 8, 7, 6, 5, 4}
	// for i := 0; i < len(numSlice); i++ {
	// 	for j := i + 1; j < len(numSlice); j++ {
	// 		if numSlice[i] > numSlice[j] {
	// 			temp := numSlice[i]
	// 			numSlice[i] = numSlice[j]
	// 			numSlice[j] = temp
	// 		}
	// 	}
	// }
	// fmt.Println(numSlice)

	//sort排序
	intList := []int{2, 4, 6, 3, 8, 2, 6, 9, 1}
	sort.Ints(intList)
	fmt.Println(intList)
	//降序写法 很麻烦
	sort.Sort(sort.Reverse(sort.IntSlice(intList)))
	fmt.Println(intList)
}
