package main

func main() {
	//golang中的变量都有一个自己对应的内存地址
	// var a = 10
	// var b = &a
	// fmt.Printf("a-%v-%T-%p\n", a, a, &a)
	// fmt.Printf("b-%v-%T-%p", b, b, &b)

	//写法错误
	//因为这里a只定义了 但指针类型是引用类型
	//这里仅定义无分配内存地址 所以默认为nil
	//下边再对nil解引用 故报错
	// var a *int
	// *a = 100
	// fmt.Println(a)

	//new make都能用来分配内存
	//make 用于初始化 slice、map 和 channel 这三种引用类型，并返回类型本身。
	//new 用于给任意类型分配内存，并返回指向该类型零值的指针。
}
