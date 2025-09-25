package main

// import "fmt"

// // interface 接口 也可以理解为标准规范
// type Usber interface {
// 	start()
// 	stop()
// }

// type Phone struct {
// 	Name string
// }

// //Phone要实现ubs接口的话 就必须实现接口内所有的方法
// func (p Phone) start() {
// 	fmt.Println(p.Name, "启动")
// }
// func (p Phone) stop() {
// 	fmt.Println(p.Name, "关机")
// }

// func main() {
// 	p := Phone{
// 		Name: "苹果手机",
// 	}
// 	// p.start()
// 	// p.stop()

// 	//golang中接口就是一个数据类型
// 	var p1 Usber
// 	p1 = p //表示手机实现Usb接口
// 	p1.start()
// 	p1.stop()
// }

// ------------------------------------
//空接口
type A interface{} //空接口 表示没有任何约束 任意类型都可以实现空接口

//空接口还可以用来表示任意类型
// func test1() {
// 	var m1 = make(map[string]interface{})
// 	m1["1"] = 123
// 	m1["2"] = "string"
// 	m1["3"] = []int{1, 2, 3}

// 	fmt.Println(m1)
// }

//如果把一个切片类型或者结构体赋值给一个空接口类型
//无法直接通过索引或者.来访问到具体内容 会发生报错
//要访问其中的切片元素，你必须先进行类型断言，将其恢复为原来的切片类型

//这里的问题就是将切片赋值给空接口以后 它不再是切片
//这时候他是一个抽象的接口类型 所以需要类型断言来恢复其切片类型
func main() {

}
