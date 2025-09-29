package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mutex sync.Mutex
var count = 0

//管道也是一种类型
//且遵循先入先出的原则

// func chtest1() {
// 	ch := make(chan int, 3)
// 	//前面是类型 后面是容量
// 	ch <- 10
// 	ch <- 20
// 	a := <-ch
// 	fmt.Println(a)
// 	fmt.Printf("值%v-容量%v-长度%v", ch, cap(ch), len(ch))
// 	//这里就可以看到 值是一个地址 说明管道是一个引用类型
// }

// func main() {
// 	chtest1()
// }
//--------------------------------------------------------

// goroutine跟channel的demo
// func putNum(intChan chan int) {
// 	for i := 2; i < 100; i++ {
// 		intChan <- i
// 	}
// 	close(intChan)
// 	wg.Done()
// }

// 不论我前面用什么方法为channel写入的数据
// 都需要在结尾添加close关闭channel
// 否则会导致的问题就是channel一直是开发状态 认为自己准备接受一个值
// 会导致接下来时候forr的时候 不是进行正常的遍历读取 而是等待新值写入

// func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
// 	for num := range intChan {
// 		var flag = true
// 		for i := 2; i < num; i++ {
// 			if num%i == 0 {
// 				flag = false
// 				break
// 			}
// 		}
// 		if flag {
// 			primeChan <- num
// 		}
// 	}
// 	//给exitChan里放入一条数据
// 	exitChan <- true
// 	wg.Done()

// }

// func printPrime(primeChan chan int) {
// 	for v := range primeChan {
// 		fmt.Println(v)
// 	}
// 	wg.Done()
// }

// func main() {
// 	intChan := make(chan int, 1000)
// 	primeChan := make(chan int, 1000)
// 	exitChan := make(chan bool, 16) //表示管道

// 	wg.Add(1)
// 	go putNum(intChan)

// 	//统计素数协程
// 	for i := 0; i < 16; i++ {
// 		wg.Add(1)
// 		go primeNum(intChan, primeChan, exitChan)
// 	}
// 	//打印素数协程
// 	wg.Add(1)
// 	go printPrime(primeChan)

// 	//判断exitChan是否存满
// 	wg.Add(1)
// 	go func() {
// 		for i := 0; i < 16; i++ {
// 			<-exitChan
// 		}
// 		//关闭primeChan
// 		close(primeChan)
// 		wg.Done()
// 	}()

// 	wg.Wait()
// 	fmt.Println("结束。。。")
// }
//--------------------------------------------------------

// 单向管道
// func test() {
// 	ch1 := make(chan<- int, 2) //只写不读
// 	ch2 := make(<-chan int, 2) //只读不写
// 	//这时候就可以用于区分读写 便于分类管理
// }
//--------------------------------------------------------

// 多路复用
// func test() {
// 	intChan := make(chan int, 10)
// 	for i := 0; i < 10; i++ {
// 		intChan <- i
// 	}

// 	stringChan := make(chan string, 5)
// 	for i := 0; i < 5; i++ {
// 		stringChan <- "hello" + fmt.Sprintf("%d", i)
// 	}

// 	//使用select来获取channel里的数据的时候不需要关闭channel
// 	for {
// 		select {
// 		case v := <-intChan:
// 			fmt.Printf("从 intChan 读取到数据%v\n", v)
// 			time.Sleep(time.Millisecond * 50)
// 		case v := <-stringChan:
// 			fmt.Printf("从 stringChan 读取到数据%v\cn", v)
// 			time.Sleep(time.Millisecond * 50)
// 		default:
// 			fmt.Printf("数据读取完毕")
// 			return
// 		}
// 	}
// }

// func main() {
// 	test()
// }
//--------------------------------------------------------

// 用互斥锁解决多个协程同时访问一个数据造成冲突问题
func test() {
	mutex.Lock()
	count++
	fmt.Println("the count is : ", count)
	time.Sleep(time.Millisecond * 50)
	mutex.Unlock()
	wg.Done()
}

func main() {
	for r := 0; r < 20; r++ {
		wg.Add(1)
		go test()
	}
	wg.Wait()
}
