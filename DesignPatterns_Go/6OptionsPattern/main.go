package main

import (
	"fmt"
	"time"
)

const (
	defaultTimeout = 10
	defaultCaching = false
)

type Connection struct {
	addr    string
	cache   bool
	timeout time.Duration
}

// 通过两个函数来实现带默认值或不带默认值
// func NewConnect(addr string) (*Connection, error) {
// 	return &Connection{
// 		addr:    addr,
// 		cache:   defaultCaching,
// 		timeout: defaultTimeout,
// 	}, nil
// }

// func NewConnectWithOptions(addr string, cache bool, timeout time.Duration) (*Connection, error) {
// 	return &Connection{
// 		addr:    addr,
// 		cache:   cache,
// 		timeout: timeout,
// 	}, nil
// }

// 相对优雅一点
// 但每次创建Connection的时候还是要创建一个ConnectionOptions
// 有点麻烦
// type ConnectionOptions struct {
// 	Caching bool
// 	Timeout time.Duration
// }

// func NewDefaultOptions() *ConnectionOptions {
// 	return &ConnectionOptions{
// 		Caching: defaultCaching,
// 		Timeout: defaultTimeout,
// 	}
// }

// func NewConnect(addr string, opts *ConnectionOptions) (*Connection, error) {
// 	return &Connection{
// 		addr:    addr,
// 		cache:   opts.Caching,
// 		timeout: opts.Timeout,
// 	}, nil
// }

type options struct {
	timeout time.Duration
	caching bool
}

type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}

func WithTimeout(t time.Duration) Option {
	return optionFunc(func(o *options) {
		o.timeout = t
	}) //这就返回了一个optionFunc函数
}

func WithCaching(cache bool) Option {
	return optionFunc(func(o *options) {
		o.caching = cache
	})
}

func NewConnect(addr string, opts ...Option) (*Connection, error) {
	options := options{
		timeout: defaultTimeout,
		caching: defaultCaching,
	}

	for _, o := range opts {
		o.apply(&options)
	}
	return &Connection{
		addr:    addr,
		cache:   options.caching,
		timeout: options.timeout,
	}, nil
}

func main() {
	fmt.Println("--- 1. 使用所有默认值 ---")
	// 不传入任何 Option，使用默认配置
	conn1, _ := NewConnect("192.168.1.1:8080")
	fmt.Printf("连接 1 (默认): 地址=%s, 缓存=%t, 超时=%s\n", conn1.addr, conn1.cache, conn1.timeout)
	// 预期: 缓存=false, 超时=10s

	fmt.Println("\n--- 2. 设置部分选项 ---")
	// 只设置超时，缓存使用默认值
	conn2, _ := NewConnect(
		"db.example.com",
		WithTimeout(5*time.Second), // 仅修改超时
	)
	fmt.Printf("连接 2 (部分): 地址=%s, 缓存=%t, 超时=%s\n", conn2.addr, conn2.cache, conn2.timeout)
	// 预期: 缓存=false, 超时=5s

	fmt.Println("\n--- 3. 设置所有选项 ---")
	// 同时设置超时和缓存
	conn3, _ := NewConnect(
		"redis.cache.net",
		WithCaching(true),
		WithTimeout(30*time.Second), // 注意：参数顺序随意
	)
	fmt.Printf("连接 3 (全部): 地址=%s, 缓存=%t, 超时=%s\n", conn3.addr, conn3.cache, conn3.timeout)
	// 预期: 缓存=true, 超时=30s
}
