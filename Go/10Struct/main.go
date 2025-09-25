package main

import (
	"encoding/json"
	"fmt"
)

//结构体基本结构跟其他语言相同
// type people struct {
// 	name string
// 	age  int
// 	sex  string
// }

//结构体首字母大写时候表示是公有的 在其他包也可以使用
//小写时表示私有

// type People struct {
// 	Name string
// 	Age  int
// 	Sex  string
// }

// func (p People) PrintInfo() {
// 	fmt.Printf("%v-%v-%v\n", p.Name, p.Age, p.Sex)
// }

// func (p *People) SetInfo(name string, age int) {
// 	p.Name = name
// 	p.Age = age
// }
//golang中结构体实例是相互独立的

//结构体允许在声明时没有字段名只有类型 称之为匿名字段
//匿名字段默认采用类型名作为字段名 且要求字段名称必须唯一

//当结构体的字段中有指针，slice，map时 需要先make分配内存 才可以使用
// type People struct {
// 	Name  string
// 	Age   int
// 	Hobby []string
// 	map1  map[string]string
// }

//关于结构体的嵌套/继承
// type Animal struct {
// 	Name string
// }

// func (a Animal) run() {
// 	fmt.Printf("%v 在运动\n", a.Name)
// }

// //子结构体
// type Dog struct {
// 	Age    int
// 	Animal //结构体嵌套 也是继承
// }

// func (d Dog) wang() {
// 	fmt.Printf("%v 在汪\n", d.Name)
// }

// func main() {
// 	var d = Dog{
// 		Age: 20,
// 		Animal: Animal{
// 			Name: "Puppy",
// 		},
// 	}
// 	d.run()
// 	d.wang()
// }
//也可以嵌套一个结构体指针
//一般用来数据共享 避免复制

// 结构体转JSON格式
// type Student struct {
// 	ID     int
// 	Gender string
// 	Name   string
// 	Sno    string
// }

// func main() {
// 	var s1 = Student{
// 		ID:     12,
// 		Gender: "男",
// 		Name:   "李四",
// 		Sno:    "0001",
// 	}

// 	fmt.Printf("%#v\n", s1)

// 	jsonByte, _ := json.Marshal(s1)
// 	jsonStr := string(jsonByte)

// 	fmt.Printf("%v\n", jsonStr)

// 	var str2 = `{"ID":12,"Gender":"男","Name":"李四","Sno":"0001"}`
// 	var s2 Student
// 	err := json.Unmarshal([]byte(str2), &s2)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Printf("%#v\n", s2)
// 	fmt.Println(s2.Name)
// }

// 结构体标签
// type Student struct {
// 	ID     int    `json:"id"`
// 	Gender string `json:"gender"`
// 	Name   string
// 	Sno    string
// }

// //如此以后 在将结构体转换成JSON格式时
// //其对应key就会发生对应改变

// func main() {
// 	var s1 = Student{
// 		ID:     12,
// 		Gender: "男",
// 		Name:   "李四",
// 		Sno:    "0001",
// 	}
// 	fmt.Printf("%#v\n", s1)
// 	jsonByte, _ := json.Marshal(s1)
// 	jsonStr := string(jsonByte)
// 	fmt.Printf("%v\n", jsonStr)
// }

type Student struct {
	Id     int
	Name   string
	Gender string
}

type Class struct {
	Title    string
	Students []Student
}

func main() {
	c := Class{
		Title:    "001班",
		Students: make([]Student, 0),
	}

	for i := 1; i < 10; i++ {
		s := Student{
			Id:     i,
			Gender: "男",
			Name:   fmt.Sprintf("stu_%v", i),
		}
		c.Students = append(c.Students, s)
	}

	// fmt.Printf("%v\n", c)

	jsonByte, _ := json.Marshal(c)
	jsonStr := string(jsonByte)
	fmt.Println(jsonStr)
}
