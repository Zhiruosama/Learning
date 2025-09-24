package main

import "fmt"

//结构体基本结构跟其他语言相同
type people struct {
	name string
	age  int
	sex  string
}

//结构体首字母大写时候表示是公有的 在其他包也可以使用
//小写时表示私有

func main() {
	var p1 = new(people)
	p1.age = 12
	p1.name = "一"
	p1.sex = "男"

	var p2 = &people{}
	p2.age = 13
	p2.name = "二"
	p2.sex = "男"

	fmt.Printf("%#v---%T\n", p1, p1)
	fmt.Printf("%#v---%T\n", p2, p2)
}
