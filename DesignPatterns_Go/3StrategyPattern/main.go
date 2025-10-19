package main

import "fmt"

// 策略模式
// 就是通过一些类来封装不同的算法
// 每一个类封装一个具体的算法（策略）

// 策略类
type IStrategy interface {
	do(int, int) int
}

// 加 策略实现
type add struct{}

func (*add) do(a, b int) int {
	return a + b
}

// 减 策略实现
type reduce struct{}

func (*reduce) do(a, b int) int {
	return a - b
}

// 策略的执行者
type Operator struct {
	strategy IStrategy
}

// 这里初始化策略类
func (operator *Operator) setStrategy(strategy IStrategy) {
	operator.strategy = strategy
}

// 调用策略中的方法
func (operator *Operator) calculate(a, b int) int {
	return operator.strategy.do(a, b)
}

func TestStratugy() {
	operator := Operator{}

	operator.setStrategy(&add{})
	result := operator.calculate(1, 2)
	fmt.Println("add", result)

	operator.setStrategy(&reduce{})
	result = operator.calculate(2, 1)
	fmt.Println("reduce", result)
}

func main() {
	TestStratugy()
}
