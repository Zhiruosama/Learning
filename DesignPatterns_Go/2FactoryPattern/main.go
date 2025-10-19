package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

type Person struct {
	Name string
	Age  int
}

type person struct {
	name string
	age  int
}

func (p Person) Greet() {
	fmt.Printf("Hi! My name is %s", p.Name)
}

func (p person) Greet() {
	fmt.Printf("Hi! My name is %s", p.name)
}

// 简单工厂模式
func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

//相比于一般的p:=&Person{}
//这样创建实例可以确保我们创建的实例具有需要的参数

// 抽象工厂模式
func AbstrctNewPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}

//这样对外就不公开内部实现 但又可以调用提供的功能
// func main() {
// 	AbstrctNewPerson("名字", 12).Greet()
// }

// -----example------

type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

func NewHttpClient() Doer {
	return &http.Client{}
}

type mockHTTPClient struct{}

func (*mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	res := httptest.NewRecorder()
	return res.Result(), nil
}

func NewMockHTTPClient() Doer {
	return &mockHTTPClient{}
}

// 假设这是核心业务逻辑
func FetchUserData(client Doer, url string) error {
	// 关键点：这个函数只要求传入一个实现了 Doer 接口的 client
	// 所以需要用到真实网络环境 就用NewHTTPClient()创建传入
	// mock就是mockClient 因为都满足Doer 所以根本不需要改业务逻辑
	req, _ := http.NewRequest("GET", url, nil)
	// 调用 Do 方法
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// ... 后续处理 resp ...
	return nil
}

func NewPersonFactory(age int) func(name string) Person {
	return func(name string) Person {
		return Person{
			Name: name,
			Age:  age,
		}
	}
}

func TestNPF() {
	newBaby := NewPersonFactory(1)
	baby := newBaby("json")
	fmt.Println(baby.Name)
	//这样就创建了一个带默认年龄的工厂
}
