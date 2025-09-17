package main

import "fmt"

func main() {
	//go中引入的包
	//定义的变量必须使用 否则报错
	fmt.Print("A", "B")
	fmt.Println("C")
	//Println会自动换行/空格

	var a int = 10
	fmt.Printf("a=%v a的类型为%T\n", a, a)
	b := 20
	fmt.Printf("b=%v", b)
}
