package main

import (
	"fmt"
	"reflect"
)

// 反射是一种机制
// 它能让程序在运行时探知任意变量的类型（Type）和值（Value）

// func reflectFn(x interface{}) {
// 	v := reflect.TypeOf(x)
// 	// v.Name() //类型名称
// 	// v.Kind() //类型种类
// 	fmt.Println(v)
// 	fmt.Println(v.Name(), "--", v.Kind())
// }

// func reflectVO(x interface{}) {
// 	//通过反射获取值
// 	v := reflect.ValueOf(x)
// 	fmt.Println(v)
// 	//但是获得到的v是reflect.value类型 不是int
// 	//想要获得int类型还需要进行转换
// 	a := 10 + v.Int()
// 	fmt.Println(a)
// }

// 通过反射修改一个值
// func reflectSetValue(x interface{}) {
// 	v := reflect.ValueOf(x)
// 	fmt.Println(v.Kind())

// 	fmt.Println(v.Elem().Kind()) //int64
// 	if v.Elem().Kind() == reflect.Int64 {
// 		v.Elem().SetInt(123)
// 	}
// }

// 通过反射获取struct信息
type Student struct {
	Name  string `json:"name1"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func (s Student) GetInfo() string {
	var str = fmt.Sprintf("姓名:%v 年龄:%v 成绩:%v", s.Name, s.Age, s.Score)
	return str
}

func (s *Student) SetInfo(name string, age int, score int) {
	s.Name = name
	s.Age = age
	s.Score = score
}

func PrintStructField(s interface{}) {
	//判断是不是结构体类型
	t := reflect.TypeOf(s)
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("不是结构体")
		return
	}

	//获取结构体信息
	field0 := t.Field(0)
	fmt.Printf("%#v\n", field0)
	fmt.Println("字段名称: ", field0.Name)
	fmt.Println("字段类型: ", field0.Type)
}

// 反射修改结构体数据
func reflectChangeStruct(s interface{}) {
	// t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	a := v.MethodByName("GetInfo").Call(nil)
	fmt.Println(a)

	var params []reflect.Value
	params = append(params, reflect.ValueOf("老刘"))
	params = append(params, reflect.ValueOf(31))
	params = append(params, reflect.ValueOf(99))
	v.MethodByName("SetInfo").Call(params)

	b := v.MethodByName("GetInfo").Call(nil)
	fmt.Println(b)
}

// 修改结构体属性
func reflectChange(s interface{}) {
	// t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	name := v.Elem().FieldByName("Name")
	name.SetString("老李")

	age := v.Elem().FieldByName("Age")
	age.SetInt(56)

}

func main() {
	s := Student{
		Name:  "老王",
		Age:   23,
		Score: 25,
	}
	reflectChange(&s)
	fmt.Println(s)
}
