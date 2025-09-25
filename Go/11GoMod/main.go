package main

// import (
// 	"fmt"
// 	"main/calc"
// 	"main/tools"
// )

// func main() {
// 	fmt.Println(calc.Add(10, 2))
// 	fmt.Println(tools.Mul(4, 5))
// 	tools.PrintInfo()
// }

// // main包中init优先于main函数
// func init() {
// 	fmt.Println("main init...")
// }

// // 如果引入的包里有init的话 被引用的包中的init会优先被调用

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	price, err := decimal.NewFromString("136.02")
	if err != nil {
		panic(err)
	}
	fmt.Println(price)
}
