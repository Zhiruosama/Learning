package main

import "fmt"
import "math"
import "runtime"

func fortest(a int) int {
	//go的for特殊在 没有() 只有{}
	sum := 0
	for i := 0; i < a; i++ {
		sum += i
	}
	return sum
}

func whiletest(a int) int {
	//go没有while
	//go只有for 当for只有一个条件的时候就是while
	//
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	return sum
}

func pow(x, n, lim float64) float64 {
	//这里的判断条件是v<lim
	//在条件表达式前可以执行一条语句
	//该语句声明的变量作用域仅在if内s
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("第 %d 次迭代，z 的值为: %g\n", i+1, z)
	}
	return z
}

func switchtest() {
	fmt.Println("当前Go语言运行环境为：")
	switch os := runtime.GOOS; os {
	case "drawn":
		fmt.Println("macOs")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Printf("%s\n", os)
	}
}

func main() {
	// fmt.Println(whiletest(10))
	// fmt.Println(pow(2, 3, 10))
	// fmt.Println(Sqrt(2))
	// switchtest()

	//defer 语句会将函数推迟到外层函数返回之后执行。
	defer fmt.Println("world")
	fmt.Println("hello")
}
