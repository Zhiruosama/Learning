package main

func main() {
	//if的使用跟别的语言几乎无异
	//Golang中的if跟for最特别的地方就是它提倡条件不加括号
	// flag := true
	// if flag {
	// 	fmt.Println("flag=true")
	// }
	//if中的大括号不能省略

	//for示例
	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("%v\n", i)
	// }
	//Golang中没有while循环
	//可以用for来代替while

	//遍历还有for range提供使用
	//for range可以遍历切片 哈希表等多种数据结构
	// var str = "你好golang"
	// for key, value := range str {
	// 	fmt.Println("key=", key, "value=", value)
	// }
	// var arr = []string{"a", "b", "c"}
	// for k, v := range arr {
	// 	fmt.Printf("key=%v,value=%v\n", k, v)
	// }

	//switch基本用法
	// var ext = ".css"
	// switch ext {
	// case ".html":
	// 	fmt.Println("test1")
	// case ".css":
	// 	fmt.Println(".css")
	// default:
	// 	fmt.Println("nah")
	// }
	//golang中的break可以不写

	//golang默认是有break的 就是不写也会在满足某一个条件以后结束往下的穿透
	//如果需要穿透可以使用fallthrough关键字
	// age := 30
	// switch {
	// case age < 20:
	// 	fmt.Println("小年轻")
	// case age > 24 && age < 60:
	// 	fmt.Println("不再年轻")
	// 	fallthrough
	// case age > 60 && age < 70:
	// 	fmt.Println("穿透测试")
	// default:
	// 	fmt.Println("老了")
	// }
	//是的fallthrough只会穿一层 需要继续往下穿就需要写多个fallthrough

	//关于break 除了常用用法 还有在多重for循环中配合label 来标记要跳出的循环
	// for i := 0; i < 2; i++ {
	// label:
	// 	for j := 0; j < 10; j++ {
	// 		if j == 3 {
	// 			break label
	// 		}
	// 		fmt.Printf("i=%v j=%v\n", i, j)
	// 	}
	// }
	// continue同样也可以结合label来使用
}
