package main

import (
	"fmt"
	"time"
)

//Go中的函数跟别的语言基本无异
//基本格式为func 函数名(参数)类型{}
//当不写返回类型的时候 就是无返回值
// func Ts(x ...int) {
// 	fmt.Printf("%v--%T", x, x)
// }
//...表示可变参数 参数数量不确定

//return还可以一次返回多个值
// func calc(x, y int) (int, int) {
// 	sum := x + y
// 	sub := x - y
// 	return sum, sub
// }

// 案例
// func mapSort(Map map[string]string) string {
// 	var sliceKey []string

// 	for k, _ := range Map {
// 		sliceKey = append(sliceKey, k)
// 	}

// 	sort.Strings(sliceKey)

// 	var str string
// 	for _, v := range sliceKey {
// 		str += fmt.Sprintf("%v=>%v", v, Map[v])
// 	}
// 	return str
// }

//回调函数
// func calc(x, y int, cb func(int, int) int) int {
// 	return cb(x, y)
// }

//匿名函数 即无名字的函数
// func main() {
// 	func() {
// 		fmt.Println("test...")
// 	}()
// }

//闭包的基本写法：
//函数里嵌套一个函数 最后返回这个函数
// func adder() func() int {
// 	var i = 10
// 	return func() int {
// 		i++
// 		return i
// 	}
// }

//闭包可以让一个变量常驻内存
//可以让这个变量不污染全局
// func main() {
// 	var a = adder()
// 	fmt.Println(a())
// 	fmt.Println(a())
// 	fmt.Println(a())
// }

//defer用于延迟执行函数
// func fn1() {
// 	fmt.Println("开始")

// 	defer func() {
// 		fmt.Println("test1")
// 		fmt.Println("test2")
// 	}()

// 	fmt.Println("结束")
// }

//defer有注册顺序跟执行顺序
//注册顺序就是代码出现的顺序
//defer有个需要注意的点就是在执行时 其输入的参数需要确定
//也就是说在执行前 会按照其代码的顺序 执行参数中的方法=>确定其参数

//Go中没有异常机制
//Go使用panic/recover模式处理错误
//panic可以在任何地方引发 recover只有在defer调用的函数中有效
// func fn1() {
// 	defer func() {
// 		err := recover()
// 		if err != nil {
// 			fmt.Println("err:", err)
// 		}
// 	}()
// 	panic("抛出一个异常")
// }

// func fn2(a int, b int) int {
// 	defer func() {
// 		err := recover()
// 		if err != nil {
// 			fmt.Println("err:", err)
// 		}
// 	}()
// 	return a / b
// }

// func timetest() {
// 	timeObj := time.Now()
// 	//12小时制
// 	fmt.Println(timeObj.Format("2006-01-02 03:04:05"))
// 	//24小时制
// 	fmt.Println(timeObj.Format("2006-01-02 15:04:05"))
// }

// func timetest2() {
// 	//时间戳转日期字符串
// 	timeObj := time.Now()
// 	unixtime := timeObj.Unix()
// 	fmt.Println("当前时间戳: ", unixtime)
// 	timeObj_now := time.Unix(unixtime, 0)
// 	fmt.Println(timeObj_now.Format("2006-01-02 03:04:05"))
// }

// func timetest3() {
// 	//日期字符串转换时间戳
// 	var str = "2025-09-23 15:12:58"
// 	var tmp = "2006-01-02 15:04:05"

// 	timeObj, _ := time.ParseInLocation(tmp, str, time.Local)
// 	fmt.Println(timeObj)
// 	fmt.Println(timeObj.Unix())
// }

func ticker() {
	ticker := time.NewTicker(time.Second)

	for T := range ticker.C {
		fmt.Println(T)
	}
}

func main() {
	ticker()
}
