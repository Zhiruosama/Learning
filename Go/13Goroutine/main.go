package main

import (
	"sync"
)

// Goroutine 是一种轻量级的并发执行单元，让你可以同时运行多个任务
// Channel 则是用于这些 Goroutine 之间安全地进行数据通信和同步的“管道”

var wg sync.WaitGroup

// func test() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println("timeF() goroutine-", i)
// 		time.Sleep(time.Millisecond * 200)
// 	}
// 	wg.Done()
// }

//	func main() {
//		wg.Add(1) //协程计数器+1
//		go test() //开启一个协程
//		for i := 0; i < 10; i++ {
//			fmt.Println("main() goroutine-", i)
//			time.Sleep(time.Millisecond * 50)
//		}
//		//当协程运行的速度比较慢时会出现问题
//		//主进程会先行完成任务 这时候协程也会被随即结束 即没运行完就结束了
//		//这时候可以使用sync.WaitGroup来解决
//		wg.Wait()
//	}
//
// ----------------------------------------------------
// func test(num int) {
// 	defer wg.Done()
// 	for i := 1; i <= 5; i++ {
// 		fmt.Printf("第%v个协程打印的第%v条数据\n", num, i)
// 	}
// }

// func main() {
// 	for i := 1; i <= 6; i++ {
// 		wg.Add(1)
// 		go test(i)
// 	}
// 	wg.Wait()
// }
// ----------------------------------------------------

// func test(n int) {
// 	for num := (n-1)*30 + 1; num <= n*30; num++ {
// 		if num > 1 {
// 			flag := true
// 			for i := 2; i < num; i++ {
// 				if num%i == 0 {
// 					flag = false
// 					break
// 				}
// 			}
// 			if flag {
// 				fmt.Println(num, "是素数")
// 			}
// 		}
// 	}
// 	wg.Done()
// }

// func main() {
// 	for i := 0; i < 4; i++ {
// 		wg.Add(1)
// 		go test(i)
// 	}
// 	wg.Wait()
// 	fmt.Println("主线程结束")
// }
// ----------------------------------------------------
