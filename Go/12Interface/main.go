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

func main() {

}
