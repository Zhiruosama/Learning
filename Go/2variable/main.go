package main

import "fmt"

func getUserinfo() (string, int) {
	return "test", 10
}

func main() {
	//go变量首个字符不能是数字
	// var 1a string 错误

	//一次性进行多个声明
	// var (
	// 	name string
	// 	age  int
	// )
	// name = "murane"
	// age = 20
	// fmt.Println(name, age)

	//短变量声明 只能声明局部变量
	// test := "测试"
	// fmt.Printf(test)

	// var username, age = getUserinfo()
	// fmt.Println(username, age)

	// var username, _ = getUserinfo() //使用匿名变量
	// fmt.Println(username)

	//iota是计数器 初始为0 往下自增 只有定义多个常量时候有用
	const (
		n1 = iota
		n2
		n3
		n4
	)
	fmt.Println(n1, n2, n3, n4)
}
