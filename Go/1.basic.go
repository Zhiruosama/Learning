package main

//Go程序由包组成
//我们需要通过import来导入各种包
//类似于cpp中的include
import "fmt"
import "math"

//函数的用法跟一般的语言一样
//需要注意的是类型在变量名后面
func add(x int,y int) int{
	return x+y
}
//Go 的返回值可被命名，它们会被视作定义在函数顶部的变量
//没有参数的 return 语句会直接返回已命名的返回值
func test(sum int)(x,y int){
	x=sum+1;
	y=sum*2;
	return 
}

func main(){
	fmt.Println("Hello World")

	//在 Go 中，如果一个名字以大写字母开头，那么它就是已导出的
	//常见的就有Pi这种
	fmt.Println(math.Pi)

	fmt.Println(test(123))

	//Go的基本类型有
	// 	bool
	// string
	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte // uint8 的别名
	// rune // int32 的别名
	//      // 表示一个 Unicode 码位
	// float32 float64
	// complex64 complex128

	// 没有明确初始化的变量声明会被赋予对应类型的 零值。
	// 零值是：
	// 数值类型为 0，
	// 布尔类型为 false，
	// 字符串为 ""（空字符串）

	
}

