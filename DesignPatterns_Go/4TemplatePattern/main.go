package main

import "fmt"

// 模板模式就是将一个类中能够公共使用的方法放置在抽象类中实现
// 将不能公共使用的方法作为抽象方法强制子类去实现
// 这样就做到了将一个类作为一个模板，让开发者去填充需要填充的地方

type Cooker interface {
	fire()
	cooke()
	outfire()
}

// 类似一个抽象类
type CookMenu struct {
}

func (CookMenu) fire() {
	fmt.Println("开火")
}

//做菜 交给具体的子类实现
func (CookMenu) cooke() {
}

func (CookMenu) outfire() {
	fmt.Println("关火")
}

//封装具体步骤
func doCook(cook Cooker) {
	cook.fire()
	cook.cooke()
	cook.outfire()
}

type XiHongShi struct {
	CookMenu
}

func (*XiHongShi) cooke() {
	fmt.Println("西红柿")
}

type JiDan struct {
	CookMenu
}

func (*JiDan) cooke() {
	fmt.Println("鸡蛋")
}

func TestTemplate() {
	xihongshi := &XiHongShi{}
	doCook(xihongshi)

	fmt.Println("==>>下一道菜")

	jidan := &JiDan{}
	doCook(jidan)
}

func main() {
	TestTemplate()
}
