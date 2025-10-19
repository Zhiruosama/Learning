package main

import (
	"sync"
)

type singleton struct {
}

var ins *singleton
var mu sync.Mutex
var once sync.Once

// func GetInxOr() *singleton {
// 	//如果不加锁 导致的问题就是
// 	//在多线程的时候有可能会创建多个实例
// 	if ins == nil {
// 		mu.Lock()
// 		if ins == nil {
// 			ins = &singleton{}
// 		}
// 		mu.Unlock()
// 	}
// 	return ins
// }

// 更优雅的写法
func GetInxOr() *singleton {
	once.Do(func() {
		ins = &singleton{}
	})
	return ins
}

// once.Do可以确保ins实例全局只被创建一次
// once.Do函数还可以确保当同时有多个创建动作时
// 只有一个创建动作在被执行
